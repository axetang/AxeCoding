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
 func (a *IPAddr) Network() string
 func (a *IPAddr) String() string
 ------------------------------------------------------------------------------------
 **Description:
 IPAddr represents the address of an IP end point.
 Network returns the address's network name, "ip".
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
