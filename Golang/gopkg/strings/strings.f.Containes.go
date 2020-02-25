/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Contains,strings.ContainsAny,strings.ContainsRune
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
------------------------------------------------------------------------------------
**Description:
Contains reports whether substr is within s.
ContainsAny reports whether any Unicode code points in chars are within s.
ContainsRune reports whether the Unicode code point r is within s.
------------------------------------------------------------------------------------
**要点总结:
1. Contains函数在s中查找substr是否存在；
2. ContainsAny函数在s中查找字符串参数chars中任何一个Unicode字符是否存在；
3. ContainsRune函数在s中查找rune参数r是否存在。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "01234飞哥78传说9"
	fmt.Println(strings.Contains(s, "01"))
	fmt.Println(strings.Contains(s, "4飞"))
	fmt.Println(strings.ContainsAny(s, "a哥"))
	fmt.Println(strings.ContainsAny(s, "糖"))
	fmt.Println(strings.ContainsRune(s, '说'))
	fmt.Println(strings.ContainsRune(s, '糖'))
}
