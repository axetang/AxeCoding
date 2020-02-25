/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Int64Var
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int64Var(p *int64, name string, value int64, usage string)
------------------------------------------------------------------------------------
 **Description:
 Int64Var defines an int64 flag with specified name, default value, and usage
 string. The argument p points to an int64 variable in which to store the
 value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Int64Var函数定义一个int64 flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *int64是一个指向定义的int64 flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i64 int64

func init() {
	flag.Int64Var(&i64, "i", 3, "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("i64:", i64)
}
