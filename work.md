 run command:

 /Users/chenhui/github/mediasoup/worker/out/Release/mediasoup-worker --logLevel=warn --logTag=info --logTag=ice --logTag=dtls --logTag=rtp --logTag=srtp --logTag=rtcp --logTag=rtx --logTag=bwe --logTag=score --logTag=simulcast --logTag=svc --logTag=sctp --rtcMinPort=40000 --rtcMaxPort=49999



mediasoup-demo-server:INFO protoo connection request [roomId:jkopkbso, peerId:i0qxuk9f, address:100.100.68.249, origin:https://100.100.68.249:3000] +42s
  mediasoup-demo-server:INFO creating a new Room [roomId:jkopkbso] +0ms
  mediasoup-demo-server:INFO:Room create() [roomId:jkopkbso] +3s
  mediasoup:Worker createRouter() +42s
  mediasoup:Channel request() [method:worker.createRouter, id:2] +3s
  mediasoup:Channel request succeeded [method:worker.createRouter, id:2] +0ms
  mediasoup:Router constructor() +3s
  mediasoup:Router createAudioLevelObserver() +0ms
  mediasoup:Channel request() [method:router.createAudioLevelObserver, id:3] +1ms
  mediasoup:Channel request succeeded [method:router.createAudioLevelObserver, id:3] +0ms
  mediasoup:RtpObserver constructor() +3s
  mediasoup:Router createDirectTransport() +0ms
  mediasoup:Channel request() [method:router.createDirectTransport, id:4] +0ms
  mediasoup:Channel request succeeded [method:router.createDirectTransport, id:4] +1ms
  mediasoup:Transport constructor() +3s
  mediasoup:DirectTransport constructor() +2m
  mediasoup:Transport produceData() +1ms
  mediasoup:Channel request() [method:transport.produceData, id:5] +1ms
  mediasoup:Channel request succeeded [method:transport.produceData, id:5] +0ms
  mediasoup:DataProducer constructor() +3s
  mediasoup-demo-server:Room protoo Peer "request" event [method:getRouterRtpCapabilities, peerId:i0qxuk9f] +3s
  mediasoup-demo-server:Room protoo Peer "request" event [method:createWebRtcTransport, peerId:i0qxuk9f] +79ms
  mediasoup:Router createWebRtcTransport() +161ms
  mediasoup:Channel request() [method:router.createWebRtcTransport, id:6] +160ms
  mediasoup:Channel request succeeded [method:router.createWebRtcTransport, id:6] +1ms
  mediasoup:Transport constructor() +161ms
  mediasoup:WebRtcTransport constructor() +2m
  mediasoup:Transport pause() +0ms
  mediasoup:Channel request() [method:transport.enableTraceEvent, id:7] +0ms
  mediasoup:Channel request succeeded [method:transport.enableTraceEvent, id:7] +1ms
  mediasoup:Transport setMaxIncomingBitrate() [bitrate:1500000] +1ms
  mediasoup:Channel request() [method:transport.setMaxIncomingBitrate, id:8] +0ms
  mediasoup:Channel request succeeded [method:transport.setMaxIncomingBitrate, id:8] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:createWebRtcTransport, peerId:i0qxuk9f] +16ms
  mediasoup:Router createWebRtcTransport() +16ms
  mediasoup:Channel request() [method:router.createWebRtcTransport, id:9] +13ms
  mediasoup:Channel request succeeded [method:router.createWebRtcTransport, id:9] +1ms
  mediasoup:Transport constructor() +16ms
  mediasoup:WebRtcTransport constructor() +17ms
  mediasoup:Transport pause() +0ms
  mediasoup:Channel request() [method:transport.enableTraceEvent, id:10] +1ms
  mediasoup:Channel request succeeded [method:transport.enableTraceEvent, id:10] +0ms
  mediasoup:Transport setMaxIncomingBitrate() [bitrate:1500000] +1ms
  mediasoup:Channel request() [method:transport.setMaxIncomingBitrate, id:11] +1ms
  mediasoup:Channel request succeeded [method:transport.setMaxIncomingBitrate, id:11] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:join, peerId:i0qxuk9f] +25ms
  mediasoup:Transport consumeData() +22ms
  mediasoup:Channel request() [method:transport.consumeData, id:12] +22ms
  mediasoup:Channel request succeeded [method:transport.consumeData, id:12] +1ms
  mediasoup:DataConsumer constructor() +3s
  mediasoup-demo-server:Room protoo Peer "request" event [method:connectWebRtcTransport, peerId:i0qxuk9f] +194ms
  mediasoup:WebRtcTransport connect() +216ms
  mediasoup:Channel request() [method:transport.connect, id:13] +192ms
  mediasoup:Channel request succeeded [method:transport.connect, id:13] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:connectWebRtcTransport, peerId:i0qxuk9f] +18ms
  mediasoup:WebRtcTransport connect() +19ms
  mediasoup:Channel request() [method:transport.connect, id:14] +18ms
  mediasoup:Channel request succeeded [method:transport.connect, id:14] +1ms
  mediasoup-demo-server:Room WebRtcTransport "sctpstatechange" event [sctpState:connecting] +2ms
  mediasoup-demo-server:Room WebRtcTransport "sctpstatechange" event [sctpState:connected] +2ms
  mediasoup-demo-server:Room WebRtcTransport "sctpstatechange" event [sctpState:connecting] +65ms
  mediasoup:WARN:Channel [pid:40361] RTC::Transport::ReceiveRtpPacket() | no suitable Producer for received RTP packet [ssrc:910142085, payloadType:111] +2m
  mediasoup:WARN:Channel [pid:40361] RTC::Transport::ReceiveRtpPacket() | no suitable Producer for received RTP packet [ssrc:910142085, payloadType:111] +1ms
  mediasoup:WARN:Channel [pid:40361] RTC::Transport::ReceiveRtpPacket() | no suitable Producer for received RTP packet [ssrc:910142085, payloadType:111] +4ms
  mediasoup:WARN:Channel [pid:40361] RTC::Transport::ReceiveRtpPacket() | no suitable Producer for received RTP packet [ssrc:910142085, payloadType:111] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:produce, peerId:i0qxuk9f] +22ms
  mediasoup:Transport produce() +303ms
  mediasoup:Channel request() [method:transport.produce, id:15] +91ms
  mediasoup:Channel request succeeded [method:transport.produce, id:15] +4ms
  mediasoup:Producer constructor() +3s
  mediasoup:RtpObserver addProducer() +511ms
  mediasoup:Channel request() [method:rtpObserver.addProducer, id:16] +1ms
  mediasoup:Channel request succeeded [method:rtpObserver.addProducer, id:16] +1ms
  mediasoup-demo-server:Room WebRtcTransport "sctpstatechange" event [sctpState:connected] +77ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:produceData, peerId:i0qxuk9f] +16ms
  mediasoup:Transport produceData() +92ms
  mediasoup:Channel request() [method:transport.produceData, id:17] +86ms
  mediasoup:Channel request succeeded [method:transport.produceData, id:17] +1ms
  mediasoup:DataProducer constructor() +598ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:produceData, peerId:i0qxuk9f] +26ms
  mediasoup:Transport produceData() +26ms
  mediasoup:Channel request() [method:transport.produceData, id:18] +24ms
  mediasoup:Channel request succeeded [method:transport.produceData, id:18] +2ms
  mediasoup:DataProducer constructor() +27ms
  mediasoup:Transport consumeData() +4ms
  mediasoup:Channel request() [method:transport.consumeData, id:19] +2ms
  mediasoup:Channel request succeeded [method:transport.consumeData, id:19] +1ms
  mediasoup:DataConsumer constructor() +424ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +10ms
  mediasoup:WebRtcTransport getStats() +219ms
  mediasoup:Channel request() [method:transport.getStats, id:20] +5ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:20] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +16ms
  mediasoup:WebRtcTransport getStats() +17ms
  mediasoup:Channel request() [method:transport.getStats, id:21] +16ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:21] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +10ms
  mediasoup:Producer getStats() +149ms
  mediasoup:Channel request() [method:producer.getStats, id:22] +9ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:22] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +7ms
  mediasoup:DataProducer getStats() +41ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:23] +6ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:23] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +7ms
  mediasoup:DataProducer getStats() +6ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:24] +5ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:24] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +5ms
  mediasoup:DataConsumer getStats() +50ms
  mediasoup:Channel request() [method:dataConsumer.getStats, id:25] +3ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:25] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:produce, peerId:i0qxuk9f] +417ms
  mediasoup:Transport produce() +469ms
  mediasoup:Channel request() [method:transport.produce, id:26] +418ms
  mediasoup:Channel request succeeded [method:transport.produce, id:26] +1ms
  mediasoup:Producer constructor() +439ms
  mediasoup:WARN:Channel [pid:40361] RTC::Transport::ReceiveRtpPacket() | no suitable Producer for received RTP packet [ssrc:561191062, payloadType:97] +667ms
  mediasoup-demo-server:Room producer "videoorientationchange" event [producerId:91194d61-480d-4658-8056-7543a488df03, videoOrientation:{ camera: false, flip: false, rotation: 0 }] +78ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +2s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:27] +2s
  mediasoup:Channel request succeeded [method:transport.getStats, id:27] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +34ms
  mediasoup:WebRtcTransport getStats() +34ms
  mediasoup:Channel request() [method:transport.getStats, id:28] +31ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:28] +7ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +29ms
  mediasoup:Producer getStats() +2s
  mediasoup:Channel request() [method:producer.getStats, id:29] +21ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:29] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +22ms
  mediasoup:Producer getStats() +22ms
  mediasoup:Channel request() [method:producer.getStats, id:30] +19ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:30] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +5ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:31] +3ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:31] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +6ms
  mediasoup:DataProducer getStats() +6ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:32] +5ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:32] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +8ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:33] +4ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:33] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:34] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:34] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +22ms
  mediasoup:WebRtcTransport getStats() +22ms
  mediasoup:Channel request() [method:transport.getStats, id:35] +22ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:35] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +26ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:36] +22ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:36] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +17ms
  mediasoup:Producer getStats() +17ms
  mediasoup:Channel request() [method:producer.getStats, id:37] +17ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:37] +6ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +24ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:38] +17ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:38] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +21ms
  mediasoup:DataProducer getStats() +22ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:39] +20ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:39] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +14ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:40] +11ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:40] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:41] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:41] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +8ms
  mediasoup:WebRtcTransport getStats() +7ms
  mediasoup:Channel request() [method:transport.getStats, id:42] +7ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:42] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +12ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:43] +10ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:43] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +18ms
  mediasoup:Producer getStats() +18ms
  mediasoup:Channel request() [method:producer.getStats, id:44] +15ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:44] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +4ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:45] +3ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:45] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +2ms
  mediasoup:DataProducer getStats() +2ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:46] +2ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:46] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +2ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:47] +2ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:47] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:48] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:48] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +9ms
  mediasoup:WebRtcTransport getStats() +9ms
  mediasoup:Channel request() [method:transport.getStats, id:49] +8ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:49] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +24ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:50] +24ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:50] +4ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +9ms
  mediasoup:Producer getStats() +9ms
  mediasoup:Channel request() [method:producer.getStats, id:51] +4ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:51] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +3ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:52] +4ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:52] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +4ms
  mediasoup:DataProducer getStats() +3ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:53] +3ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:53] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +4ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:54] +5ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:54] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:55] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:55] +0ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +26ms
  mediasoup:WebRtcTransport getStats() +26ms
  mediasoup:Channel request() [method:transport.getStats, id:56] +25ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:56] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +19ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:57] +18ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:57] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +19ms
  mediasoup:Producer getStats() +20ms
  mediasoup:Channel request() [method:producer.getStats, id:58] +19ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:58] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +36ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:59] +33ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:59] +6ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +16ms
  mediasoup:DataProducer getStats() +16ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:60] +10ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:60] +9ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +38ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:61] +29ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:61] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:62] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:62] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +24ms
  mediasoup:WebRtcTransport getStats() +25ms
  mediasoup:Channel request() [method:transport.getStats, id:63] +22ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:63] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +9ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:64] +7ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:64] +5ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +12ms
  mediasoup:Producer getStats() +12ms
  mediasoup:Channel request() [method:producer.getStats, id:65] +7ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:65] +2ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +5ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:66] +3ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:66] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +7ms
  mediasoup:DataProducer getStats() +7ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:67] +6ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:67] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataConsumerStats, peerId:i0qxuk9f] +3ms
  mediasoup:DataConsumer getStats() +3s
  mediasoup:Channel request() [method:dataConsumer.getStats, id:68] +2ms
  mediasoup:Channel request succeeded [method:dataConsumer.getStats, id:68] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +3s
  mediasoup:WebRtcTransport getStats() +3s
  mediasoup:Channel request() [method:transport.getStats, id:69] +3s
  mediasoup:Channel request succeeded [method:transport.getStats, id:69] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getTransportStats, peerId:i0qxuk9f] +10ms
  mediasoup:WebRtcTransport getStats() +10ms
  mediasoup:Channel request() [method:transport.getStats, id:70] +9ms
  mediasoup:Channel request succeeded [method:transport.getStats, id:70] +3ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +7ms
  mediasoup:Producer getStats() +3s
  mediasoup:Channel request() [method:producer.getStats, id:71] +4ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:71] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getProducerStats, peerId:i0qxuk9f] +28ms
  mediasoup:Producer getStats() +28ms
  mediasoup:Channel request() [method:producer.getStats, id:72] +27ms
  mediasoup:Channel request succeeded [method:producer.getStats, id:72] +10ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +33ms
  mediasoup:DataProducer getStats() +3s
  mediasoup:Channel request() [method:dataProducer.getStats, id:73] +28ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:73] +1ms
  mediasoup-demo-server:Room protoo Peer "request" event [method:getDataProducerStats, peerId:i0qxuk9f] +56ms
  mediasoup:DataProducer getStats() +51ms
  mediasoup:Channel request() [method:dataProducer.getStats, id:74] +50ms
  mediasoup:Channel request succeeded [method:dataProducer.getStats, id:74] +5ms
  protoo-server:ERROR:WebSocketTransport connection "error" event [conn:WSS:[100.100.68.249]:63142, error:Error: read ECONNRESET
    at TLSWrap.onStreamRead (node:internal/stream_base_commons:211:20) {
  errno: -54,
  code: 'ECONNRESET',
  syscall: 'read'
}] +23s
  mediasoup-demo-server:Room protoo Peer "close" event [peerId:i0qxuk9f] +26ms
  mediasoup:Transport close() +19s
  mediasoup:Channel request() [method:transport.close, id:75] +21ms
  mediasoup:Producer transportClosed() +116ms
  mediasoup:Producer transportClosed() +0ms
  mediasoup:DataProducer transportClosed() +27ms
  mediasoup:DataProducer transportClosed() +0ms
  mediasoup:Transport close() +1ms
  mediasoup:Channel request() [method:transport.close, id:76] +1ms
  mediasoup:DataConsumer transportClosed() +3s
  mediasoup-demo-server:INFO:Room last Peer in the room left, closing the room [roomId:jkopkbso] +20s
  mediasoup-demo-server:Room close() +5ms
  mediasoup:Router close() +20s
  mediasoup:Channel request() [method:router.close, id:77] +4ms
  mediasoup:Transport routerClosed() +5ms
  mediasoup:DataProducer transportClosed() +5ms
  mediasoup:DataConsumer transportClosed() +1ms
  mediasoup:RtpObserver routerClosed() +20s
  mediasoup:Channel request succeeded [method:transport.close, id:75] +1ms
  mediasoup:Channel request succeeded [method:transport.close, id:76] +2ms
  mediasoup:Channel request succeeded [method:router.close, id:77] +1ms