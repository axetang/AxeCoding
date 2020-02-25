/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Float64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Float64(name string, value float64, usage string) *float64
------------------------------------------------------------------------------------
 **Description:
 Float64 defines a float64 flag with specified name, default value, and usage
 string. The return value is the address of a float64 variable that stores
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Float64函数定义一个float64 flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 返回值是指向这个float64 flag的指针*float64。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var f64 = flag.Float64("pi", 3.14, "input pi")

func main() {
	flag.Parse()
	fmt.Println("f64:", *f64)
}
