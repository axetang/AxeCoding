/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.UintVar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func UintVar(p *uint, name string, value uint, usage string)
------------------------------------------------------------------------------------
 **Description:
 UintVar defines a uint flag with specified name, default value, and usage
 string. The argument p points to a uint variable in which to store the
 value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. UintVar函数定义一个uint flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *uint是一个指向定义的uint flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i uint

func init() {
	flag.UintVar(&i, "i", 3, "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("i:", i)
}
