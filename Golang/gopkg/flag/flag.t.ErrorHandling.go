/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.ErrorHandling
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type ErrorHandling int
 const (
    ContinueOnError ErrorHandling = iota // Return a descriptive error.
    ExitOnError                          // Call os.Exit(2).
    PanicOnError                         // Call panic with a descriptive error.
 )
 ------------------------------------------------------------------------------------
 **Description:
 ErrorHandling defines how FlagSet.Parse behaves if the parse fails.
 These constants cause FlagSet.Parse to behave as described if the parse fails.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ErrorHandling定义如何处理flag解析错误;
 2. 这些定义的常数使得FlagSet.Parse方法在出现解析错误时产生对应的行动：
   1） ContinueOnError:返回一个描述性的错误；
   2） ExitOnError：调用os.Exit(2)，退出;
   3) PanicOnError：调用panic，同时提供一个描述性的错误信息。
*************************************************************************************/
/*
package main

import (
	"flag"
	"fmt"
)

var ui = flag.Uint("pi", 3, "input pi")

func main() {

	flag.Parse()
	fmt.Println("Uint:", *ui)
}
*/
