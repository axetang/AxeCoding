/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.ToLower,bytes.ToLowerSpecial
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ToLower(s []byte) []byte
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
------------------------------------------------------------------------------------
**Description:
ToLower treats s as UTF-8-encoded bytes and returns a copy with all the Unicode
letters mapped to their lower case.
ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode
letters mapped to their lower case, giving priority to the special casing rules.
------------------------------------------------------------------------------------
**要点总结:
1. ToLower返回s的一个副本，并把其中所有的Unicode字符转换为小写。
2. ToLowerSpecial返回s的一个副本，并把其中的所有Unicode字符都根据_case指定的规则转换成小写。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	println("--------ToLower------------")
	s := []byte("Hello, World!")
	fmt.Println(string(bytes.ToLower(s)))
	fmt.Println(string(s))
	println("--------ToLowerSpecial------------")
	s = []byte("Hello, world!")
	fmt.Println(string(bytes.ToLowerSpecial(unicode.AzeriCase, s)))
	fmt.Println(string(s))
}
