/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.RuneStart
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func RuneStart(b byte) bool
 ------------------------------------------------------------------------------------
 **Description:
 RuneStart reports whether the byte could be the first byte of an encoded, possibly
 invalid rune. Second and subsequent bytes always have the top two bits set to 10.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneStart函数判断参数b byte是否是一个合法的rune字符编码的第一个字节；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
}
