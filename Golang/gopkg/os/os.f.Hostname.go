/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Hostname
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Hostname() (name string, err error)
  ------------------------------------------------------------------------------------
 **Description:
 Hostname returns the host name reported by the kernel.
 ------------------------------------------------------------------------------------
 **要点总结:
 Hostname函数返回计算机名称。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The hostname is: %s\n", name)
}
