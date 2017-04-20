# gogarp

A tool for sending IPv4 GARP(Gratuitous ARP) packet.

```
go get github.com/higebu/gogarp
sudo setcap cap_net_raw=ep $GOPATH/bin/gogarp
gogarp eth0
```
