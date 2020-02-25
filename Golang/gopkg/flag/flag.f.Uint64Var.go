/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Uint64Var
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Uint64Var(p *uint64, name string, value uint64, usage string)
------------------------------------------------------------------------------------
 **Description:
 Uint64Var defines a uint64 flag with specified name, default value, and
 usage string. The argument p points to a uint64 variable in which to store
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Uint64Var函数定义一个uint64 flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *uint64是一个指向定义的uint64 flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i uint64

func init() {
	flag.Uint64Var(&i, "i", 3, "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("i:", i)
}
