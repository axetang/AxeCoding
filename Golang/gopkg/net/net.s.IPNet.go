/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPNet
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type IPNet struct {
     IP   IP     // network number
     Mask IPMask // network mask
 }
 func (n *IPNet) Contains(ip IP) bool
 func (n *IPNet) Network() string
 func (n *IPNet) String() string
 ------------------------------------------------------------------------------------
 **Description:
 An IPNet represents an IP network.
 Contains reports whether the network includes ip.
 Network returns the address's network name, "ip+net".
 String returns the CIDR notation of n like "192.0.2.1/24" or "2001:db8::/48" as
 defined in RFC 4632 and RFC 4291. If the mask is not in the canonical form, it
 returns the string which consists of an IP address, followed by a slash character
 and a mask expressed as hexadecimal form with no punctuation like
 "198.51.100.1/c000ff00".
 ------------------------------------------------------------------------------------
 **要点总结:
 0. IPNet代表一个IP网络，包含两个成员：IP net.IP和Mask net.IPMask;
 1. Contains返回此IP网络是否包含该IP;
 2. Network返回此IP网络名称  "ip+net";
 3. String返回IP网络的字符串形式。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip, ipnet, err := net.ParseCIDR("198.168.0.0/16")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("IP:%#v", ip)         //返回 IP:[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc6, 0xa8, 0x0, 0x0}
	fmt.Printf("IPNet:%#v\n", ipnet) //返回 IPNet:&net.IPNet{IP:[]byte{0xc6, 0xa8, 0x0, 0x0}, Mask:[]byte{0xff, 0xff, 0x0, 0x0}}

	ip1 := net.ParseIP("198.168.0.12")
	fmt.Println("Contains:", ipnet.Contains(ip1))
	ip2 := net.ParseIP("198.168.1.20")
	fmt.Println("Contains:", ipnet.Contains(ip2))

	fmt.Println("Network:", ipnet.Network())
	fmt.Println(ipnet.String())
}
