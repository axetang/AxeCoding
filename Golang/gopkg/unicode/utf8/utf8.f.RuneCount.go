/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.RuneCount
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func RuneCount(p []byte) int
 ------------------------------------------------------------------------------------
 **Description:
 RuneCount returns the number of runes in p. Erroneous and short encodings are
 treated as single runes of width 1 byte.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneCount方法返回参数p []byte中的rune字符数
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
}
