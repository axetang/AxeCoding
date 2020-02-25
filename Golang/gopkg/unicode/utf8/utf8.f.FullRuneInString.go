/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.FullRuneInString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FullRuneInString(s string) bool
 ------------------------------------------------------------------------------------
 **Description:
 FullRuneInString is like FullRune but its input is a string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FullRuneInString和FullRune方法类似，但它的输入是一个字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "i世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]), len(str))
}
