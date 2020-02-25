/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.SPrint
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
func Sprint(a ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 Sprint formats using the default formats for its operands and returns the resulting
 string. Spaces are added between operands when neither is a string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sprint使用其操作数的默认格式进行格式化并返回其结果字符串;
 2. 当两个连续的操作数均不为字符串时，它们之间就会添加空格;如下面程序，即使操作数中含有字符串，也会
 自动添加空格；
 2. 查阅标准库源码，Sprint定义和实现如下
 func Sprint(a ...interface{}) string {
	p := newPrinter()
	p.doPrint(a)
	s := string(p.buf)
	p.free()
	return s
 }
*************************************************************************************/
package main

import (
	"fmt"
)

func main() {
	i := 0
	b := false
	f := 3.14
	s := "abc is abcdef"
	ss := fmt.Sprintln(i, b, f, s)
	sss := fmt.Sprintln(i, s, b, f)
	fmt.Println(ss)
	fmt.Println(sss)
}
