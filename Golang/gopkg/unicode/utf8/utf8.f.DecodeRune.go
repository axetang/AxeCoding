/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.DecodeRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DecodeRune(p []byte) (r rune, size int)
 ------------------------------------------------------------------------------------
 **Description:
 DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and its width
 in bytes. If p is empty it returns (RuneError, 0). Otherwise, if the encoding is
 invalid, it returns (RuneError, 1). Both are impossible results for correct,
 non-empty UTF-8.
 An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of
 range, or is not the shortest possible UTF-8 encoding for the value. No other
 validation is performed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. DecodeRune函数解码输入参数p []byte的第一个UTF-8字符并返回该rune字符及其size；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
	}
}
