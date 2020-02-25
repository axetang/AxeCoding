/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Parse
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Parse()
------------------------------------------------------------------------------------
 **Description:
 Parse parses the command-line flags from os.Args[1:]. Must be called after
 all flags are defined and before flags are accessed by the program.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Parse函数从os.Args[1:]中解析命令行flag；
 2. 此函数必须在所有flags被定义后以及在flags可以被程序访问前执行。
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
