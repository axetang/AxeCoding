/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.RuneCountInString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func RuneCountInString(s string) (n int)
 ------------------------------------------------------------------------------------
 **Description:
 RuneCountInString is like RuneCount but its input is a string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneCountInString方法和RuneCount方法类似，但它的输入参数是一个字符串s。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))
}
