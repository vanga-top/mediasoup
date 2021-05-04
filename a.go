// Copyright 2017 QINIU. All rights reserved.

package souplpcsvc

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"

	"github.com/someonegg/msgpump/msgpeer"
	"go.uber.org/atomic"
	xlog "qiniupkg.com/x/xlog.v7"

	"github.com/qbox/pili-pegasus/cmd/pili-roomworker/internal/config"
	"github.com/qbox/pili-pegasus/cmd/pili-roomworker/internal/model"
	"github.com/qbox/pili-pegasus/internal/common/service/netstring"
	"github.com/qbox/pili-pegasus/internal/common/utility"
	"github.com/qbox/pili-pegasus/internal/protodef/souplpc"
)

type logSoupd struct {
	xl *xlog.Logger
}

func (l *logSoupd) Write(p []byte) (n int, err error) {
	l.xl.Infof("[SOUPD_STD_LOG] %s", p)
	return len(p), nil
}

func newLogSoupd(ID string) (l *logSoupd) {
	l = &logSoupd{
		xl: xlog.New(ID),
	}
	return l
}

type Client struct {
	conf     *config.SoupdConfig
	mset     *model.ManagerSet
	p        *msgpeer.Peer
	cmd      *exec.Cmd
	xl       *xlog.Logger
	exitChan chan error
	online   *atomic.Bool
	soupdID  string
	stdLog   *logSoupd
}

// ID implement for model.Soupd interface
func (c *Client) ID() string {
	return c.soupdID
}

// Online implement for model.Soupd interface
func (c *Client) Online() bool {
	return c.online.Load()
}

// Close implement for model.Soupd interface
func (c *Client) Close() error {
	if c.p != nil {
		c.p.Stop()
	}

	if c.cmd != nil {
		c.cmd.Process.Signal(syscall.SIGTERM)
	}

	return nil
}

// Request implement for model.Soupd interface
func (c *Client) Request(ctx context.Context, t string, v model.SoupMessage) ([]byte, error) {
	m, err := v.Marshal()
	if err != nil {
		return nil, err
	}

	if c.online.Load() {
		if t != souplpc.MT_WebRtcTransport_GetStats && t != souplpc.MT_Producer_GetStats && t != souplpc.MT_Consumer_GetStats {
			c.xl.Infof("request to %v, %v=%v", c.soupdID, t, string(m))
		}
		res, err := c.p.Do(ctx, t, m)
		if t != souplpc.MT_WebRtcTransport_GetStats && t != souplpc.MT_Producer_GetStats && t != souplpc.MT_Consumer_GetStats {
			c.xl.Infof("response from %v, %v=%v, err: %v", c.soupdID, t, string(res), err)
		}

		return res, err
	}

	return nil, model.ErrSoupdUnavailable
}

// Process implement for mesgpeer.Handler interface
func (c *Client) Process(ctx context.Context, t string, m msgpeer.Request, w msgpeer.ResponseWriter) {
	defer func() {
		if e := recover(); e != nil {
			const size = 16 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.xl.Error("process panic: ", e, fmt.Sprintf("\n%s", buf))
		}
	}()

	c.xl.Errorf("No support request from %v, %v=%v", c.soupdID, t, string(m))
}

// OnNotify implement for mesgpeer.Handler interface
func (c *Client) OnNotify(ctx context.Context, t string, m msgpeer.Notify) {
	defer func() {
		if e := recover(); e != nil {
			const size = 16 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.xl.Error("notify panic: ", e, fmt.Sprintf("\n%s", buf))
		}
	}()

	if t == "logger" {
		switch m[0] {
		case 'D':
			c.xl.Debug("[SOUPD_DEBUG]", string(m[1:]))
		case 'W':
			c.xl.Warn("[SOUPD_WARN]", string(m[1:]))
		case 'E':
			c.xl.Error("[SOUPD_ERROR]", string(m[1:]))
		default:
			c.xl.Error(string(m))
		}

		return
	}

	c.xl.Infof("notify from %v, %v=%v", c.soupdID, t, string(m))

	switch t {
	case souplpc.MT_Running:
		c.onRunning(m)
	case souplpc.MT_Score:
		c.onScore(m)
	case souplpc.MT_VideoOrientationChange:
		c.onVideoOrientationChange(m)
	case souplpc.MT_IceSelectedTupleChange:
		c.onIceSelectedTupleChange(m)
	case souplpc.MT_IceStateChange:
		c.onIceStateChange(m)
	case souplpc.MT_DtlsStateChange:
		c.onDtlsStateChange(m)
	case souplpc.MT_ProducerClose:
		c.onProducerClose(m)
	case souplpc.MT_LayersChange:
		c.onLayersChange(m)
	default:
		c.xl.Errorf("no support notify from %v, %v=%v", c.soupdID, t, string(m))
	}
}

func (c *Client) dialMediasoup(args []string, h msgpeer.Handler) (*msgpeer.Peer, *exec.Cmd, error) {
	channelReadPairs, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	channelReadFather := os.NewFile(uintptr(channelReadPairs[0]), "channelReadFather")
	channelReadChild := os.NewFile(uintptr(channelReadPairs[1]), "channelReadChild")

	channelWritePairs, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	channelWriteFather := os.NewFile(uintptr(channelWritePairs[0]), "channelWriteFather")
	channelWriteChild := os.NewFile(uintptr(channelWritePairs[1]), "channelWriteChild")

	payloadChannelReadPairs, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	payloadChannelReadFather := os.NewFile(uintptr(payloadChannelReadPairs[0]), "payloadChannelReadFather")
	payloadChannelReadChild := os.NewFile(uintptr(payloadChannelReadPairs[1]), "payloadChannelReadChild")
	payloadChannelWritePairs, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	payloadChannelWriteFather := os.NewFile(uintptr(payloadChannelWritePairs[0]), "payloadChannelWriteFather")
	payloadChannelWriteChild := os.NewFile(uintptr(payloadChannelWritePairs[1]), "payloadChannelWriteChild")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.ExtraFiles = []*os.File{channelWriteChild, channelReadChild, payloadChannelWriteChild, payloadChannelReadChild}
	cmd.Env = append(os.Environ(), "MEDIASOUP_VERSION=3")
	cmd.Stderr = c.stdLog
	connR, err := net.FileConn(channelReadFather)
	if err != nil {
		return nil, nil, err
	}
	connW, err := net.FileConn(channelWriteFather)
	if err != nil {
		return nil, nil, err
	}
	rw := netstring.NetstringMRW(bufio.NewReader(connR), bufio.NewWriter(connW))
	peer := msgpeer.NewPeer(rw, h, c.conf.PumpWriteQueue)
	peer.Start(nil)
	cmd.Start()
	go func() {
		c.exitChan <- cmd.Wait()
		channelReadFather.Close()
		channelWriteFather.Close()
		payloadChannelReadFather.Close()
		payloadChannelWriteFather.Close()
	}()

	return peer, cmd, nil
}

func Start(conf *config.SoupdConfig, mset *model.ManagerSet, soupdID string, minPort, maxPort int) *Client {
	c := &Client{
		conf:     conf,
		mset:     mset,
		xl:       xlog.New(utility.NewREQID()),
		exitChan: make(chan error, 1),
		online:   atomic.NewBool(false),
		soupdID:  soupdID,
		stdLog:   newLogSoupd(soupdID),
	}

	go func() {
		args := []string{
			conf.MediasoupPath,
			soupdID,
			"--logLevel=warn",
			"--logTag=info",
			"--logTag=ice",
			"--logTag=dtls",
			"--logTag=rtp",
			"--logTag=srtp",
			"--logTag=rtcp",
			"--logTag=rtx",
			"--logTag=bwe",
			"--logTag=score",
			"--logTag=simulcast",
			"--logTag=svc",
			"--logTag=sctp",
			"--logTag=message",
			"--rtcMinPort=" + strconv.Itoa(minPort),
			"--rtcMaxPort=" + strconv.Itoa(maxPort),
			"--dtlsCertificateFile=" + conf.DtlsCertificateFile,
			"--dtlsPrivateKeyFile=" + conf.DtlsPrivateKeyFile,
		}
		p, cmd, err := c.dialMediasoup(args, c)
		if err != nil {
			c.xl.Errorf("dialMediasoup failed %v, error: %v", c.soupdID, err)
		}
		c.p = p
		c.cmd = cmd
		c.xl.Infof("dialMediasoup succeed %v", c.soupdID)
		c.online.Store(true)
		c.mset.SoupdM.Connect(c)
		select {
		case err = <-c.exitChan:
			c.xl.Errorf("mediasoup exit %v, error: %v", c.soupdID, err.Error())
			c.online.Store(false)
			c.mset.RoomM.OnSoupdDisconnected(c.soupdID)
			c.mset.SoupdM.Disconnect(c.soupdID)
		case <-c.p.StopD():
			c.xl.Errorf("mediasoup stopped %v, error: %v", c.soupdID, c.p.Error())
			c.online.Store(false)
			c.mset.RoomM.OnSoupdDisconnected(c.soupdID)
			c.mset.SoupdM.Disconnect(c.soupdID)
		}
	}()
	return c
}

func (c *Client) onRunning(m []byte) {
	var notify souplpc.WorkerRunningNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}
	// TODO
}

func (c *Client) onScore(m []byte) {
	var notify souplpc.WorkerRunningNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}
	// TODO
}

func (c *Client) onVideoOrientationChange(m []byte) {
	var notify souplpc.VideoOrientationChangeNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}
	// TODO
}

func (c *Client) onIceStateChange(m []byte) {
	var notify souplpc.IceStateChangeNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}
	room := c.mset.RoomM.GetRoomByTransportID(notify.TargetID)
	if room == nil {
		c.xl.Errorf("No room hold transportID: %v", notify.TargetID)
		return
	}
	err = room.OnSoupdNotifyIceStateChange(&notify)
	if err != nil {
		c.xl.Errorf("notify process failed: %v", err)
	}
	return
}

func (c *Client) onIceSelectedTupleChange(m []byte) {
	var notify souplpc.IceSelectedTupleChangeNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}

	room := c.mset.RoomM.GetRoomByTransportID(notify.TargetID)
	if room == nil {
		c.xl.Errorf("No room hold transportID: %v", notify.TargetID)
		return
	}

	err = room.OnSoupdNotifyIceSelectedTupleChange(&notify)
	if err != nil {
		c.xl.Errorf("notify process failed: %v", err)
	}

	return
}

func (c *Client) onDtlsStateChange(m []byte) {
	var notify souplpc.DtlsStateChangeNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}

	room := c.mset.RoomM.GetRoomByTransportID(notify.TargetID)
	if room == nil {
		c.xl.Errorf("No room hold transportID: %v", notify.TargetID)
		return
	}

	err = room.OnSoupdNotifyDtlsStateChange(&notify)
	if err != nil {
		c.xl.Errorf("notify process failed: %v", err)
	}
	return
}

func (c *Client) onProducerClose(m []byte) {
	var notify souplpc.ProducerCloseNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}

	room := c.mset.RoomM.GetRoomByConsumerID(notify.TargetID)
	if room == nil {
		c.xl.Errorf("No room hold consumerID: %v", notify.TargetID)
		return
	}

	err = room.OnSoupdNotifyProducerClose(&notify)
	if err != nil {
		c.xl.Errorf("notify process failed: %v", err)
	}
	return
}

func (c *Client) onLayersChange(m []byte) {
	var notify souplpc.LayersChangeNotify
	err := notify.Unmarshal(m)
	if err != nil {
		return
	}

	room := c.mset.RoomM.GetRoomByConsumerID(notify.TargetID)
	if room == nil {
		c.xl.Errorf("No room hold ConsumerID: %v", notify.TargetID)
		return
	}

	err = room.OnSoupdNotifyLayersChange(&notify)
	if err != nil {
		c.xl.Errorf("notify process failed: %v", err)
	}
	return
}
