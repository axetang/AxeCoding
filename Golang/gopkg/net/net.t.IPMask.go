/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPMask
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type IPMask []byte
 func CIDRMask(ones, bits int) IPMask
 func IPv4Mask(a, b, c, d byte) IPMask
 func (m IPMask) Size() (ones, bits int)
 func (m IPMask) String() string
 ------------------------------------------------------------------------------------
 **Description:
 An IP mask is an IP address.
 CIDRMask returns an IPMask consisting of `ones' 1 bits followed by 0s up to a total
 length of `bits' bits. For a mask of this form, CIDRMask is the inverse of IPMask.Size.
 IPv4Mask returns the IP mask (in 4-byte form) of the IPv4 mask a.b.c.d.
 Size returns the number of leading ones and total bits in the mask. If the mask is
 not in the canonical form--ones followed by zeros--then Size returns 0, 0.
 String returns the hexadecimal form of m, with no punctuation.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. IPMask代表一个IP地址的掩码;
 1. CIDRMask返回一个IPMask类型值，该返回值总共有bits个字位，其中前ones个字位都是1，其余字位都是0.
 2. IPv4Mask返回一个4字节格式的IPv4掩码a.b.c.d;
 3. String返回m的十六进制格式，没有标点;
 4. Size返回m的前导的1字位数和总字位数。如果m不是规范的子网掩码（字位：/^1+0+$/），将返会(0, 0);
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
