/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Fields,strings.FieldsFunc
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
------------------------------------------------------------------------------------
**Description:
Fields splits the string s around each instance of one or more consecutive white
space characters, as defined by unicode.IsSpace, returning a slice of substrings of
s or an empty slice if s contains only white space.
FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
and returns an array of slices of s. If all code points in s satisfy f(c) or the
string is empty, an empty slice is returned. FieldsFunc makes no guarantees about
the order in which it calls f(c). If f does not return consistent results for a
given c, FieldsFunc may crash.
------------------------------------------------------------------------------------
**要点总结:
1. Fields函数使用一个或连续的多个空格来切分参数字符串s，生成新的字符串切片返回；如果s只有空格，
则返回空切片；
2. FieldsFunc函数功能和Fields类似，但通过函数参数f的返回值来判断是否为用于切分的分隔符。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz  飞哥  "))
	println()

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3.飞哥，baz4..", f))
	println()
}
