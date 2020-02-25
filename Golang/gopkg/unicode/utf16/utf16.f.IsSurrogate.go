/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf16
 **Element: utf16.IsSurrogate
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsSurrogate(r rune) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsSurrogate reports whether the specified Unicode code point can appear in a
 surrogate pair.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IsSurrogate判断参数r rune是否为代理区字符。两个代理区字符可以用来组合成一个utf16编码。
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
