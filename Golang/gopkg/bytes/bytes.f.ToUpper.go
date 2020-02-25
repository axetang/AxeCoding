/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.ToUpper,bytes.ToUpperSpecial
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ToUpper(s []byte) []byte
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
------------------------------------------------------------------------------------
**Description:
ToUpper treats s as UTF-8-encoded bytes and returns a copy with all the Unicode
letters within it mapped to their upper case.
ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the
Unicode letters mapped to their upper case, giving priority to the special casing
rules.
------------------------------------------------------------------------------------
**要点总结:
1. ToUpper返回s的一个副本，并把其中所有的Unicode字符转换为大写。
2. ToUpperSpecial返回s的一个副本，并把其中的所有Unicode字符都根据_case指定的规则转换成大写。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	println("------ToUpper-------")
	s := []byte("hello, world!")
	fmt.Println(string(bytes.ToUpper(s)))
	fmt.Println(string(s))

	println("------ToUpperSpecial-------")
	s = []byte("Hello, world!")
	fmt.Println(string(bytes.ToUpperSpecial(unicode.AzeriCase, s)))
	fmt.Println(string(s))
}
