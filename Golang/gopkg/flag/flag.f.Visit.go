/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Visit
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Visit(fn func(*Flag))
------------------------------------------------------------------------------------
 **Description:
 Visit visits the command-line flags in lexicographical order, calling fn
 for each. It visits only those flags that have been set.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Visit函数按照字典顺序遍历flag，并且对每个标签调用fn；
 2. Visit函数只遍历解析时进行设置了的标签。
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

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.Visit.go -s=1 -bool=T a
	//执行结果：
	//1 true
	//NFlag(): 2
	//Name:bool, Usage:Description of flag,Value:true,DefValue:false
	//Name:s, Usage:Description,Value:1,DefValue:
	flag.Visit(PrintFlag)
}
