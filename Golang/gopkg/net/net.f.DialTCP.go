/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.DialTCP
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
 ------------------------------------------------------------------------------------
 **Description:
 DialTCP acts like Dial for TCP networks.
 The network must be a TCP network name; see func Dial for details.
 If laddr is nil, a local address is automatically chosen. If the IP field of raddr
 is nil or an unspecified IP address, the local system is assumed.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main
