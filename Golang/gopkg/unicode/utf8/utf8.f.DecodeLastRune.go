/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.DecodeLastRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DecodeLastRune(p []byte) (r rune, size int)
 ------------------------------------------------------------------------------------
 **Description:
 DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and its
 width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if the encoding
 is invalid, it returns (RuneError, 1). Both are impossible results for correct,
 non-empty UTF-8.
 An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of
 range, or is not the shortest possible UTF-8 encoding for the value. No other
 validation is performed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. DecodeLastRune函数解码输入参数p []byte最后一个UTF-8编码字符并返回；
 2. 如果p为空则返回错误（RuneError，0）；
 3. 如果解码错误则返回（RuneError，1）。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[:len(b)-size]
	}

	b = []byte{}
	r, size := utf8.DecodeLastRune(b)
	fmt.Printf("%c %#U %v\n", r, r, size)
}
