/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Args
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Args() []string
------------------------------------------------------------------------------------
 **Description:
 Args returns the non-flag command-line arguments.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Args函数返回non-flag命令行参数的字符串切片[]string。
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
	//执行如下命令验证程序执行结果：
	//go run flag.f.Args.go -s=1 -bool=T a b c
	//运行结果：
	//1 true
	//Args(): [a b c]

}
