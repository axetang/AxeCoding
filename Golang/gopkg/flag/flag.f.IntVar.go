/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.IntVar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IntVar(p *int, name string, value int, usage string)
------------------------------------------------------------------------------------
 **Description:
 IntVar defines an int flag with specified name, default value, and usage
 string. The argument p points to an int variable in which to store the
 value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. IntVar函数定义一个int flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *int是一个指向定义的int flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i int

func init() {
	flag.IntVar(&i, "i", 3, "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("i:", i)
}
