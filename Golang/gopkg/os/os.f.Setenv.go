/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Setenv
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Setenv(key, value string) error
 ------------------------------------------------------------------------------------
 **Description:
 Setenv sets the value of the environment variable named by the key. It returns an
 error, if any.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Setenv方法使用参数key和value设置一个环境变量；
 2. 只在当前进程有效。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Path of go is: %s\n", os.Getenv("go"))
	if err := os.Setenv("go", "/path/to/go"); err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Now Path of go is: %s\n", os.Getenv("go"))

}
