/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.DecodeLastRuneInString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DecodeLastRuneInString(s string) (r rune, size int)
 ------------------------------------------------------------------------------------
 **Description:
 DecodeLastRuneInString is like DecodeLastRune but its input is a string. If s is
 empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns
 (RuneError, 1). Both are impossible results for correct, non-empty UTF-8.
 An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of
 range, or is not the shortest possible UTF-8 encoding for the value. No other
 validation is performed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. DecodeLastRuneInString和DecodeLastRune方法类似，但它的输入是一个字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[:len(str)-size]
	}
}
