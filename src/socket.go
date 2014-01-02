package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "192.168.0.191"
	ipaddr := net.ParseIP(ip)
	fmt.Println(ipaddr.To4())
}
