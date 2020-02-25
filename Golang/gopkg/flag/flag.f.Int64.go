/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Int64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int64(name string, value int64, usage string) *int64
------------------------------------------------------------------------------------
 **Description:
 Int64 defines an int64 flag with specified name, default value, and usage
 string. The return value is the address of an int64 variable that stores
 the value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
1. Int函数定义一个Int64 flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 返回值是指向这个int64 flag的指针*int64。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i64 = flag.Int("i64", 3, "input pi")

func main() {
	flag.Parse()
	fmt.Println("int64:", *i64)
}
