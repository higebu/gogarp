package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/mdlayher/arp"
)

func main() {
	flag.Parse()
	ifName := flag.Arg(0)
	if ifName == "" {
		fmt.Println("please specify interface name. ex: gogarp eth0")
		os.Exit(1)
	}

	ifi, err := net.InterfaceByName(ifName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c, err := arp.Dial(ifi)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	addrs, err := ifi.Addrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, addr := range addrs {
		ip, _, _ := net.ParseCIDR(addr.String())
		if ip.To4() != nil {
			p, err := arp.NewPacket(arp.OperationRequest, ifi.HardwareAddr, ip, net.HardwareAddr{0, 0, 0, 0, 0, 0}, ip)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = c.WriteTo(p, net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
