/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf16
 **Element: utf16.EncodeRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func EncodeRune(r rune) (r1, r2 rune)
 ------------------------------------------------------------------------------------
 **Description:
 EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune. If the rune
 is not a valid Unicode code point or does not need encoding, EncodeRune returns
 U+FFFD, U+FFFD.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. EncodeRune将字符r编码成UTF-16代理对。参数r是要编码的字符；
 2. 如果r<0x10000，则无需编码，其UTF-16序列就是其自身
 	r1：编码后的 UTF-16 代理对的高位码元
	r2：编码后的 UTF-16 代理对的低位码元
 3. 如果r不是有效的Unicode字符，或者是代理区字符,或者无需编码，则返回 U+FFFD, U+FFFD。
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
