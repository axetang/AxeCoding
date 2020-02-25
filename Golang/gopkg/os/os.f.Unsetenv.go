/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Unsetenv
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Unsetenv(key string) error
 ------------------------------------------------------------------------------------
 **Description:
 Unsetenv unsets a single environment variable.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnsetEnv函数删除key string指定的环境变量值；
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

	os.Unsetenv("go")
	fmt.Printf("Now Path of go is: %s\n", os.Getenv("go"))

}
