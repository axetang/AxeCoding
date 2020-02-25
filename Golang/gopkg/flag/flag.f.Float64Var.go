/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Float64Var
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Float64Var(p *float64, name string, value float64, usage string)
------------------------------------------------------------------------------------
 **Description:
 Float64Var defines a float64 flag with specified name, default value, and
 usage string. The argument p points to a float64 variable in which to store
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Float64Var函数定义一个float64 flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 函数第一个参数p *float64是一个指向定义的float64 flag的指针，存储命令行输入的值
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var f64 float64

func init() {
	flag.Float64Var(&f64, "pi", 3.14, "input pi")
}

func main() {
	flag.Parse()
	fmt.Println("f64:", f64)
}
