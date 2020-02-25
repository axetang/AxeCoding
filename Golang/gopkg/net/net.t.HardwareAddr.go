/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.HardwareAddr
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type HardwareAddr []byte
 func (a HardwareAddr) String() string
 ------------------------------------------------------------------------------------
 **Description:
 A HardwareAddr represents a physical hardware address.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. HardwareAddr类型代表一个硬件地址（MAC地址）;
 1. String方法返回HardwareAddr类型的字符串表达。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	hw, err := net.ParseMAC("01:23:45:67:89:ab")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", hw) //返回 []byte{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}
	fmt.Println("hw.String()", hw.String())
}
