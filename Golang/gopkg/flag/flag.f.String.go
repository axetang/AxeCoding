/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.String
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func String(name string, value string, usage string) *string
------------------------------------------------------------------------------------
 **Description:
 String defines a string flag with specified name, default value, and usage
 string. The return value is the address of a string variable that stores
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. String函数定义一个string flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数返回一个指向定义的string flag的指针，存储命令行输入的flag的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var s = flag.String("s", "123", "INPUT STR")

func main() {
	flag.Parse()
	fmt.Println("s:", *s)
}
