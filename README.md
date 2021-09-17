# live_websocket 中间件
websocket raw data protocol plugin for monibuca

通过Websocket传输音视频数据，使用LiveP播放器进行播放。

## 中间件地址

https://github.com/qnsoft/live_websocket.git

## 配置

可配置WS协议和WSS协议监听地址端口

```toml
[Jessica]
ListenAddr = ":8080"
CertFile = "../foo.cert"
KeyFile  = "../foo.key"
ListenAddrTLS = ":8088"
```

- 如果不设置ListenAddr和ListenAddrTLS，将共用网关的端口监听

## 协议说明

该插件提供两种格式的协议供播放器播放。

### WS-RAW

- 地址格式：ws://[HOST]/jessica/[streamPath]

- 该协议传输的是私有格式，第一个字节代表音视频，1为音频，2为视频，后面跟4个字节的时间戳，然后接上字节流（RTMP的VideoTag或AudioTag）

### WS-FLV

- 地址格式：ws://[HOST]/jessica/[streamPath].flv
- 该协议传输的flv格式的文件流
