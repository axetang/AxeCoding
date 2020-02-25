/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ListenPacket
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ListenPacket(network, address string) (PacketConn, error)
 ------------------------------------------------------------------------------------
 **Description:
 ListenPacket announces on the local network address.
 The network must be "udp", "udp4", "udp6", "unixgram", or an IP transport. The IP
 transports are "ip", "ip4", or "ip6" followed by a colon and a literal protocol
 number or a protocol name, as in "ip:1" or "ip:icmp".
 For UDP and IP networks, if the host in the address parameter is empty or a literal
 unspecified IP address, ListenPacket listens on all available IP addresses of the
 local system except multicast IP addresses. To only use IPv4, use network "udp4" or
 "ip4:proto". The address can use a host name, but this is not recommended, because
 it will create a listener for at most one of the host's IP addresses. If the port
 in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port
 number is automatically chosen. The LocalAddr method of PacketConn can be used to
 discover the chosen port.
 See func Dial for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ListenPacket函数监听本地网络地址laddr。网络类型net必须是面向数据包的网络类型："ip"、"ip4"、
 "ip6"、"udp"、"udp4"、"udp6"、或"unixgram".
*************************************************************************************/
package main

func main() {
}
