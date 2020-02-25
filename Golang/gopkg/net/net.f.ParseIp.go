/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ParseIP
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseIP(s string) IP
 ------------------------------------------------------------------------------------
 **Description:
 ParseIP parses s as an IP address, returning the result. The string s can be in
 dotted decimal ("192.0.2.1") or IPv6 ("2001:db8::68") form. If s is not a valid
 textual representation of an IP address, ParseIP returns nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseIP将s解析为IP地址，并返回该地址IP结构体。如果s不是合法的IP地址文本表示，ParseIP会返回nil;
 2. 字符串可以是小数点分隔的IPv4格式如"74.125.19.99"或IPv6格式如"2001:4860:0:2001::68"格式。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {

	ip := net.ParseIP("74.125.19.99")
	fmt.Printf("IP:%#v\n", ip)
	//返回
	//IP:[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x4a, 0x7d, 0x13, 0x63}

}
