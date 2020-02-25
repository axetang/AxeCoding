/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPv4
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IPv4(a, b, c, d byte) IP
 ------------------------------------------------------------------------------------
 **Description:
 IPv4 returns the IP address (in 16-byte form) of the IPv4 address a.b.c.d.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IPv4函数返回包含一个IPv4地址a.b.c.d的IP地址（16字节格式）;
 2. 四个参数a,b,c,d byte确定了地址内容。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.IPv4(byte(127), byte(0), byte(1), byte(1))
	fmt.Println(ip) //返回 127.0.1.1
	ip = net.IPv4(byte(255), byte(255), byte(255), byte(0))
	fmt.Println(ip) //返回
}
