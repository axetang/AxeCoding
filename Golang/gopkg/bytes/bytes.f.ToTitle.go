/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.ToTitle,bytes.TitleSpecial
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ToTitle(s []byte) []byte
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
------------------------------------------------------------------------------------
**Description:
ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode
letters mapped to their title case.
ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the
Unicode letters mapped to their title case, giving priority to the special casing
rules.
------------------------------------------------------------------------------------
**要点总结:
1. ToTitle返回s的一个副本，并把其中所有的Unicode字符转换为首字母大写；
2. ToTitleSpecial返回s的一个副本，并把其中的所有Unicode字符都根据_case指定的规则转换成首字母大写。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	println("--------ToTitle------------")
	s := []byte("helLO, wORld!")
	fmt.Println(string(bytes.ToTitle(s)))
	fmt.Println(string(s))
	println("--------ToTitleSpecial------------")
	s = []byte("helLO, wORld!")
	fmt.Println(string(bytes.ToTitleSpecial(unicode.AzeriCase, s)))
	fmt.Println(string(s))
}
