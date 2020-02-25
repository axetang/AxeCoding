/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.SPrintf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sprintf(format string, a ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 Sprintf formats according to a format specifier and returns the resulting string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查阅标准包源码，Sprintf定义和实现如下, Sprintf根据格式说明进行格式化并返回其结果字符串;
 func Sprintf(format string, a ...interface{}) string {
	p := newPrinter()
	p.doPrintf(format, a)
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
	f := -3.14
	s := "abc"
	ss := fmt.Sprintf("%d,%t,%6.2f,%s", i, b, f, s)
	fmt.Println(ss)
}
