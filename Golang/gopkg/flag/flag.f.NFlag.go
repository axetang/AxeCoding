/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.NFlag
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NFlag() int
------------------------------------------------------------------------------------
 **Description:
 NFlag returns the number of command-line flags that have been set.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. NFlag函数返回命令行已经有多少个flag参数被设置值。
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

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.NArg.go -s=1 -bool=T a b c
}
