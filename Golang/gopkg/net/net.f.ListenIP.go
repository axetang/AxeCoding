/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ListenIP
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
 ------------------------------------------------------------------------------------
 **Description:
 ListenIP acts like ListenPacket for IP networks.
 The network must be an IP network name; see func Dial for details.
 If the IP field of laddr is nil or an unspecified IP address, ListenIP listens on
 all available IP addresses of the local system except multicast IP addresses.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
