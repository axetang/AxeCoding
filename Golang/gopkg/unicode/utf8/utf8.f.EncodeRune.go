/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.EncodeRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func EncodeRune(p []byte, r rune) int
 ------------------------------------------------------------------------------------
 **Description:
 EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the
 rune. It returns the number of bytes written.
 FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a
 rune. An invalid encoding is considered a full Rune since it will convert as a
 width-1 error rune.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. EncodeRune函数将参数r rune字符采用UTF-8编码写入参数p []byte；
 2. 返回写入字节数；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}
