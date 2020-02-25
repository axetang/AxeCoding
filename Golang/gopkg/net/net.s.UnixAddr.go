/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.UnixAddr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type UnixAddr struct {
     Name string
     Net  string
 }
 UnixAddr represents the address of a Unix domain socket end point.
 func ResolveUnixAddr(network, address string) (*UnixAddr, error)
 ResolveUnixAddr returns an address of Unix domain socket end point.
 The network must be a Unix network name.
 See func Dial for a description of the network and address parameters.
 func (a *UnixAddr) Network() string
 Network returns the address's network name, "unix", "unixgram" or "unixpacket".
 func (a *UnixAddr) String() string
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
