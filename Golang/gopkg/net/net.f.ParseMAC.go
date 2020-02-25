/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ParseMAC
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseMAC(s string) (hw HardwareAddr, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address using one of the following formats:
 01:23:45:67:89:ab
 01:23:45:67:89:ab:cd:ef
 01:23:45:67:89:ab:cd:ef:00:00:01:23:45:67:89:ab:cd:ef:00:00
 01-23-45-67-89-ab
 01-23-45-67-89-ab-cd-ef
 01-23-45-67-89-ab-cd-ef-00-00-01-23-45-67-89-ab-cd-ef-00-00
 0123.4567.89ab
 0123.4567.89ab.cdef
 0123.4567.89ab.cdef.0000.0123.4567.89ab.cdef.0000
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseMAC函数使用如下格式解析一个IEEE 802 MAC-48、EUI-48或EUI-64硬件地址：
 01:23:45:67:89:ab
 01:23:45:67:89:ab:cd:ef
 01-23-45-67-89-ab
 01-23-45-67-89-ab-cd-ef
 0123.4567.89ab
 0123.4567.89ab.cdef
 2. ParseMAC函数将解析后的硬件地址以hw HardwareAddr类型的方式返回；
 3. HardwareAddr类型定义：
 	type HardwareAddr []byte
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
