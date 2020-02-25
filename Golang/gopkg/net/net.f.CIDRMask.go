/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.CIDRMask
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func CIDRMask(ones, bits int) IPMask
 ------------------------------------------------------------------------------------
 **Description:
 CIDRMask returns an IPMask consisting of `ones' 1 bits followed by 0s up to a total
 ------------------------------------------------------------------------------------
 **要点总结:
 1. CIDRMask返回一个IPMask类型值，该返回值总共有bits个字位，其中前ones个字位都是1，其余字位都是0.
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
