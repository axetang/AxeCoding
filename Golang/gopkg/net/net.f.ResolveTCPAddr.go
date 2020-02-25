/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ResolveTCPAddr
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ResolveTCPAddr(network, address string) (*TCPAddr, error)
 ------------------------------------------------------------------------------------
 **Description:
 ResolveTCPAddr returns an address of TCP end point.
 The network must be a TCP network name.
 If the host in the address parameter is not a literal IP address or the port is not a literal port number, ResolveTCPAddr resolves the address to an address of TCP end point. Otherwise, it parses the address as a pair of literal IP address and port number. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.
 See func Dial for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
