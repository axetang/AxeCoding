/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf16
 **Element: utf16.Encode
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Encode(s []rune) []uint16
 ------------------------------------------------------------------------------------
 **Description:
 Encode returns the UTF-16 encoding of the Unicode code point sequence s.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Encode将s编码成UTF-16序列并返回。
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
