/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ParseCIDR
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseCIDR(s string) (IP, *IPNet, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseCIDR parses s as a CIDR notation IP address and prefix length, like
 "192.0.2.0/24" or "2001:db8::/32", as defined in RFC 4632 and RFC 4291.
 It returns the IP address and the network implied by the IP and prefix length.
 For example, ParseCIDR("192.0.2.1/24") returns the IP address 192.0.2.1 and the
 network 192.0.2.0/24.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseCIDR将s作为一个CIDR（无类型域间路由）的IP地址和掩码字符串，如"192.168.100.1/24"或
 "2001:DB8::/48"，解析并返回IP地址和IP网络，参见RFC 4632和RFC 4291;
 2. 本函数会返回IP地址和该IP所在的网络和掩码。例如，ParseCIDR("192.168.100.1/16")会返回IP地址
 192.168.100.1和IP网络192.168.0.0/16。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ipv4addr:", ipv4Addr)
	fmt.Println("ipv4net:", ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ipv6Addr:", ipv6Addr)
	fmt.Println("ipv6Net:", ipv6Net)
}
