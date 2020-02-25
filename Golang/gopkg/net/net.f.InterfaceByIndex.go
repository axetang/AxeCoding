/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.InterfaceByIndex
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func InterfaceByIndex(index int) (*Interface, error)
 ------------------------------------------------------------------------------------
 **Description:
 InterfaceByIndex returns the interface specified by index.
 On Solaris, it returns one of the logical network interfaces sharing the logical
 data link; for more precision use InterfaceByName.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. InterfaceByIndex函数返回指定索引index int的网络接口*Interface。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	interf, err := net.InterfaceByIndex(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(interf) //返回 &{1 16436 lo  up|loopback}
}
