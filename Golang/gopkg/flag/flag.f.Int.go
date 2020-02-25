/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Int
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int(name string, value int, usage string) *int
------------------------------------------------------------------------------------
 **要点总结:
 1. Int函数定义一个Int flag；
 2. 函数参数制定了这个flag的名称，缺省值和Usage；
 3. 返回值是指向这个int flag的指针*int。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var i = flag.Int("pi", 3, "input pi")

func main() {
	flag.Parse()
	fmt.Println("int:", *i)
}
