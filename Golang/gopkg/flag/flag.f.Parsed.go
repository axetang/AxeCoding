/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Parsed
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Parsed() bool
------------------------------------------------------------------------------------
 **Description:
 Parsed reports whether the command-line flags have been parsed.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Parsed函数判断是否命令行已经被解析过。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var strFlag = flag.String("s", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func main() {
	fmt.Println("Parsed?", flag.Parsed())
	flag.Parse()
	println(*strFlag, *boolFlag)

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.NArg.go -s=1 -bool=T a b c
	fmt.Println("Parsed?", flag.Parsed())
}
