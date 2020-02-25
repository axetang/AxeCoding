/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Narg
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NArg() int
------------------------------------------------------------------------------------
 **Description:
 NArg is the number of arguments remaining after flags have been processed.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. NArg函数返回在flags处理完后，还剩下多少个non-flag命令行参数。
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

	i := flag.NArg()
	fmt.Println("NArg():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.NArg.go -s=1 -bool=T a b c
}
