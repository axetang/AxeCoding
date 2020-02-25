/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.SPrintln
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sprintln(a ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 Sprintln formats using the default formats for its operands and returns the resulting string.
 Spaces are always added between operands and a newline is appended.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查阅标准包源码，Sprintln定义和实现如下。Sprintln将操作数a切片的值按照缺省格式写入一个新的字符串
 返回，操作数之间增加空格，结尾增加换行
 func Sprintln(a ...interface{}) string {
	p := newPrinter()
	p.doPrintln(a)
	s := string(p.buf)
	p.free()
	return s
 }
 2. Sprintln调用了fmt包的私有函数newPrinter和doPrintln，具体可参见Fprintln函数的代码
*************************************************************************************/
package main

import (
	"fmt"
)

func main() {
	i := 0
	b := false
	f := 3.14
	s := "abc"
	ss := fmt.Sprintln(i, b, f, s)
	fmt.Println(ss)
}
