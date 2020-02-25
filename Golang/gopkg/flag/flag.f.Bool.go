/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Bool
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Bool(name string, value bool, usage string) *bool
------------------------------------------------------------------------------------
 **Description:
 Bool defines a bool flag with specified name, default value, and usage string.
 The return value is the address of a bool variable that stores the value of
 the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Bool函数定义一个bool flag；
 2. Bool函数参数name string表示bool flag的名称，参数value bool表示bool flag的缺省值，
 参数usage string代表参数的Usage使用说明；
 3. Bool函数返回值为指向这个bool变量的指针，所以在使用print函数时，要用*号取地址操作。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var strFlag = flag.String("s", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	ss := flag.Args()
	fmt.Println("Args():", ss)

}
