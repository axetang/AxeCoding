/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Set(name, value string) error
------------------------------------------------------------------------------------
 **Description:
 Set sets the value of the named command-line flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. 程序中调用Set函数设置命令行flag参数的值。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var strFlag = flag.String("s", "abc", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.NArg.go -s=1 -bool=T a b c
	flag.PrintDefaults()
	fmt.Println("s", *strFlag)
	fmt.Println("-----Set-----")
	flag.Set("s", "axe")
	fmt.Println("s", *strFlag)
}
