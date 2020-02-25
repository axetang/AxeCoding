/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.ValidString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ValidString(s string) bool
 ------------------------------------------------------------------------------------
 **Description:
 ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 函数ValidString判断是否字符串参数s是一个合法的UTF-8编码rune字符集合。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
}
