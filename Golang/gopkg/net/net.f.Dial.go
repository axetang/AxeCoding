/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Dial
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Dial(network, address string) (Conn, error)
 ------------------------------------------------------------------------------------
 **Description:
 Dial connects to the address on the named network.
 Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
 For TCP and UDP networks, the address has the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name. If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in RFC 4007. The functions JoinHostPort and SplitHostPort manipulate a pair of host and port in this form. When using TCP, and the host resolves to multiple IP addresses, Dial will try each IP address in order until one succeeds.
 Examples:
 Dial("tcp", "golang.org:http")
 Dial("tcp", "192.0.2.1:http")
 Dial("tcp", "198.51.100.1:80")
 Dial("udp", "[2001:db8::1]:domain")
 Dial("udp", "[fe80::1%lo0]:53")
 Dial("tcp", ":80")
 For IP networks, the network must be "ip", "ip4" or "ip6" followed by a colon and a literal protocol number or a protocol name, and the address has the form "host". The host must be a literal IP address or a literal IPv6 address with zone. It depends on each operating system how the operating system behaves with a non-well known protocol number such as "0" or "255".
 Examples:
 Dial("ip4:1", "192.0.2.1")
 Dial("ip6:ipv6-icmp", "2001:db8::1")
 Dial("ip6:58", "fe80::1%lo0")
 For TCP, UDP and IP networks, if the host is empty or a literal unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is assumed.
 For Unix networks, the address must be a file system path.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Conn接口代表通用的面向流的网络连接, 多个线程可能会同时调用同一个Conn的方法;
 1. Dial函数在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
 "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、
 "unixpacket"
 2. 对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。
 	Dial("tcp", "12.34.56.78:80")
 	Dial("tcp", "google.com:http")
 	Dial("tcp", "[2001:db8::1]:http")
 	Dial("tcp", "[fe80::1%lo0]:80")
 3. 对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值。
 	Dial("ip4:1", "127.0.0.1")
 	Dial("ip6:ospf", "::1")
 4. 对Unix网络，地址必须是文件系统路径。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conn) //返回类似 &net.TCPConn{fd:(*net.netFD)(0xf840069000)}
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())
	n, err := conn.Write([]byte("This is msg from axe!"))
	fmt.Println(n, err)
	buf := make([]byte, 10)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err = conn.Read(buf)
	fmt.Println(buf)
}
