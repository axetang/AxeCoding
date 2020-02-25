/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.VisitAll
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func VisitAll(fn func(*Flag))
------------------------------------------------------------------------------------
 **Description:
 VisitAll visits the command-line flags in lexicographical order, calling fn
 for each. It visits all flags, even those not set.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. VisitAll函数按照字典顺序遍历标签，并且对每个标签调用fn；
 2. VisitAll函数和Visit函数不同，它会遍历所有标签不管解析时有无进行设置。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

func PrintFlag(f *flag.Flag) {
	fmt.Printf("Name:%s, Usage:%s,Value:%s,DefValue:%s\n", f.Name, f.Usage, f.Value, f.DefValue)
}

var strFlag = flag.String("s", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")
var intFlag = flag.Int("i", 0, "Description of Int")

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	fmt.Println("-----Visit-----")
	flag.Visit(PrintFlag)
	fmt.Println("-----VisitAll-----")
	flag.VisitAll(PrintFlag)
	//执行如下命令验证程序执5行结果：
	//go run flag.f.Visit.go -s=1 -bool=T a
}
