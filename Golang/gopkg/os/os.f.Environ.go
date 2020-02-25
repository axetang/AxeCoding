/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Environ
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Environ() []string
 ------------------------------------------------------------------------------------
 **Description:
 Environ returns a copy of strings representing the environment, in the form
 "key=value".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Environ获取全部环境变量，采用key=value的方式写入字符串并返回。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The path is: %+v\n", os.Environ())
}
