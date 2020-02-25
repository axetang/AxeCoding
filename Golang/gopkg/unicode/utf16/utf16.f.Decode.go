/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf16
 **Element: utf16.Decode
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Decode(s []uint16) []rune
 ------------------------------------------------------------------------------------
 **Description:
 Decode returns the Unicode code point sequence represented by the UTF-16 encoding s.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Decode函数返回输入参数s []uint16表示的UTF-16编码字符切片[]rune。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	println("IsSurrogate:")
	fmt.Printf(" %t, ", utf16.IsSurrogate(0xD400)) // false
	fmt.Printf(" %t, ", utf16.IsSurrogate(0xDC00)) // true
	fmt.Printf(" %t\n", utf16.IsSurrogate(0xDFFF)) // true

	r1, r2 := utf16.EncodeRune('𠀾')
	fmt.Printf("EncodeRune: %x, %x\n", r1, r2) // d840, dc3e

	r := utf16.DecodeRune(0xD840, 0xDC3E)
	fmt.Printf("DecodeRune: %c\n", r) // d840, dc3e

	u := []uint16{'不', '会', 0xD840, 0xDC3E}
	s := utf16.Decode(u)
	fmt.Printf("%c", s) // [不 会 𠀾]
}
