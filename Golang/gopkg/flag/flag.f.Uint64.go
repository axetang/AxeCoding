/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Uint64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Uint64(name string, value uint64, usage string) *uint64
------------------------------------------------------------------------------------
 **Description:
 Uint64 defines a uint64 flag with specified name, default value, and usage
 string. The return value is the address of a uint64 variable that stores
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Uint64函数定义一个uint flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 返回值是指向这个uint64 flag的指针*uint。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var ui = flag.Uint64("pi", 3, "input pi")

func main() {
	flag.Parse()
	fmt.Println("Uint:", *ui)
}
