/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Listen
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Listen(network, address string) (Listener, error)
 ------------------------------------------------------------------------------------
 **Description:
 Listen announces on the local network address.
 The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
 For TCP networks, if the host in the address parameter is empty or a literal
 unspecified IP address, Listen listens on all available unicast and anycast IP
 addresses of the local system. To only use IPv4, use network "tcp4". The address
 can use a host name, but this is not recommended, because it will create a listener
 for at most one of the host's IP addresses. If the port in the address parameter is
 empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically
 chosen. The Addr method of Listener can be used to discover the chosen port.
 See func Dial for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
