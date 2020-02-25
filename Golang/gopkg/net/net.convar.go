/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
     IPv4len = 4
     IPv6len = 16
 )
 IP address lengths (bytes).
 var (
     IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast
     IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems
     IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers
     IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros
 )
 Well-known IPv4 addresses
常用的IPv4地址。

 var (
     IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
     IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
     IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
     IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
     IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
     IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
 )
 Well-known IPv6 addresses
 常用的IPv6地址。

 var DefaultResolver = &Resolver{}
 DefaultResolver is the resolver used by the package-level Lookup functions and by 
 Dialers without a specified Resolver.
 var (
     ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
 )
 Various errors contained in OpError.
 很多OpError类型的错误会包含本错误。
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
