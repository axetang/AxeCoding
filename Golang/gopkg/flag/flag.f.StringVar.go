/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.StringVar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func StringVar(p *string, name string, value string, usage string)
------------------------------------------------------------------------------------
 **Description:
 StringVar defines a string flag with specified name, default value, and
 usage string. The argument p points to a string variable in which to store
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. StringVar函数定义一个string flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *string是一个指向定义的string flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var s string

func init() {
	flag.StringVar(&s, "s", "123", "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("s:", s)
}
