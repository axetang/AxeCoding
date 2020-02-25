/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.UDPAddr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type UDPAddr struct {
     IP   IP
     Port int
     Zone string // IPv6 scoped addressing zone
 }
 UDPAddr represents the address of a UDP end point.
 func ResolveUDPAddr(network, address string) (*UDPAddr, error)
 ResolveUDPAddr returns an address of UDP end point.
 The network must be a UDP network name.
 If the host in the address parameter is not a literal IP address or the port is not a literal port number, ResolveUDPAddr resolves the address to an address of UDP end point. Otherwise, it parses the address as a pair of literal IP address and port number. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.
 See func Dial for a description of the network and address parameters.
 func (a *UDPAddr) Network() string
 Network returns the address's network name, "udp".
 func (a *UDPAddr) String() string
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
