/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.InterfaceByName
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func InterfaceByName(name string) (*Interface, error)
 ------------------------------------------------------------------------------------
 **Description:
 InterfaceByName returns the interface specified by name.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. InterfaceByName返回指定名字name string的网络接口*Interface。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	interf, err := net.InterfaceByName("en0")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(interf)
}
