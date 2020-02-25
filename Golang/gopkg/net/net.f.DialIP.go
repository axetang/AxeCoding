/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.DialIP
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
 ------------------------------------------------------------------------------------
 **Description:
 DialIP acts like Dial for IP networks.
 The network must be an IP network name; see func Dial for details.
 If laddr is nil, a local address is automatically chosen. If the IP field of raddr
 is nil or an unspecified IP address, the local system is assumed.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
