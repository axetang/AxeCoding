/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.SplitHostPort
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SplitHostPort(hostport string) (host, port string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 SplitHostPort splits a network address of the form "host:port", "host%zone:port",
 "[host]:port" or "[host%zone]:port" into host or host%zone and port.
 A literal IPv6 address in hostport must be enclosed in square brackets, as in
 "[::1]:80", "[::1%lo0]:80".
 See func Dial for a description of the hostport parameter, and host and port results.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SplitHostPort函数将格式为"host:port"、"[host]:port"或"[ipv6-host%zone]:port"的网络地址
 分割为host或ipv6-host%zone和port两个部分。Ipv6的文字地址或者主机名必须用方括号括起来，
 如"[::1]:80"、"[ipv6-host]:http"、"[ipv6-host%zone]:80";
 2. 参看Dial函数了解关于hostport参数的描述。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	host, port, err := net.SplitHostPort("163.com:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host, port)

	host1, port1, err1 := net.SplitHostPort("[abc:bbc]:88")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(host1, port1)

	host1, port1, err1 = net.SplitHostPort("[ipv6-host%zone]:99")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(host1, port1)
}
