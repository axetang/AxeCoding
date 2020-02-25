/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPAddr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type IPAddr struct {
     IP   IP
     Zone string // IPv6 scoped addressing zone
 }
 func ResolveIPAddr(network, address string) (*IPAddr, error)
 func (a *IPAddr) Network() string
 func (a *IPAddr) String() string
 ------------------------------------------------------------------------------------
 **Description:
 IPAddr represents the address of an IP end point.
 ResolveIPAddr returns an address of IP end point.
 The network must be an IP network name.
 If the host in the address parameter is not a literal IP address, ResolveIPAddr
 resolves the address to an address of IP end point. Otherwise, it parses the address
 as a literal IP address. The address parameter can use a host name, but this is not
 recommended, because it will return at most one of the host name's IP addresses.
 See func Dial for a description of the network and address parameters.
 Network returns the address's network name, "ip".
 ------------------------------------------------------------------------------------
 **要点总结:
i*************************************************************************************/
package main

func main() {
}
