/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Arg
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Arg(i int) string
------------------------------------------------------------------------------------
 **Description:
 Arg returns the i'th command-line argument. Arg(0) is the first remaining
 argument after flags have been processed. Arg returns an empty string if
 the requested element does not exist.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Arg函数返回flag处理后的non-flag参数，函数参数i表示顺序；
 2. Arg(0)是flag处理后的首个non-flag参数；
 3. 如果指定顺序的命令参数不存在，则返回空串。
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

	ss := flag.Arg(0)
	fmt.Println("Arg(0):", ss)
	//执行如下命令验证程序执行结果：
	//go run flag.f.Arg.go -s=1 -bool=T a b c
	//运行结果：
	//1 true
	//Arg(0): a
}
