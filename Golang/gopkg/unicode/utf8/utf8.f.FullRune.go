/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.FullRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FullRune(p []byte) bool
 ------------------------------------------------------------------------------------
 **Description:
 FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a
 rune. An invalid encoding is considered a full Rune since it will convert as a
 width-1 error rune.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FullRune函数判断参数p []byte中的字节以UTF-8编码的rune开始；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))
}
