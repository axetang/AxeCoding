/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPMask
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 func IPv4Mask(a, b, c, d byte) IPMask
 ------------------------------------------------------------------------------------
 **Description:
 IPv4Mask returns the IP mask (in 4-byte form) of the IPv4 mask a.b.c.d.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IPv4Mask返回一个4字节格式的IPv4掩码a.b.c.d;
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	//ip := net.ParseIP("192.168.100.1")
	ipm := net.IPv4Mask(255, 255, 255, 254)
	fmt.Printf("%#v,%s\n", ipm, ipm.String())
	fmt.Printf("%#d,%s\n", ipm, ipm.String())
	fmt.Printf("%#x,%s\n", ipm, ipm.String())

	ones, bits := ipm.Size()
	fmt.Println(ones, bits)

	ipm = net.CIDRMask(30, 32)
	fmt.Println(ipm.String())
}
