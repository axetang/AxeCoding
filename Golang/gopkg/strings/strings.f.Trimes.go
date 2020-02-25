/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Trim,strings.TrimFunc,strings.TrimLeftFunc,strings.TrimPrefix,
strings.TrimRight,strings.TrimRightFunc,strings.TrimSpace,string.TrimSuffix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Trim(s string, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s string, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s string, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
------------------------------------------------------------------------------------
**Description:
Trim returns a slice of the string s with all leading and trailing Unicode code
points contained in cutset removed.
TrimFunc returns a slice of the string s with all leading and trailing Unicode code
points c satisfying f(c) removed.
TrimLeft returns a slice of the string s with all leading Unicode code points
contained in cutset removed.
To remove a prefix, use TrimPrefix instead.
TrimLeftFunc returns a slice of the string s with all leading Unicode code points c
satisfying f(c) removed.
TrimPrefix returns s without the provided leading prefix string. If s doesn't start
with prefix, s is returned unchanged.
TrimRight returns a slice of the string s, with all trailing Unicode code points
contained in cutset removed.
To remove a suffix, use TrimSuffix instead.
TrimRightFunc returns a slice of the string s with all trailing Unicode code points
c satisfying f(c) removed.
TrimSpace returns a slice of the string s, with all leading and trailing white space
removed, as defined by Unicode.
TrimSuffix returns s without the provided trailing suffix string. If s doesn't end
with suffix, s is returned unchanged.
------------------------------------------------------------------------------------
**要点总结:
1. Trim函数将字符串s的头部和尾部的cutset子字符串中字符剪去后返回；
2. TrimFunc函数将字符串s头部和尾部按照函数参数f定义剪去后返回，遇到第一个f(c)为false时就停止一侧剪切；
3. TrimLeft函数将字符串s头部的cutset子字符串中的字符剪去后返回；
4. TrimLeftFunc函数将字符串s头部按照函数参数f定义剪去后返回；
5. TrimRight函数将字符串s尾部的cutset子字符串剪去后返回；
6. TrimRightFunc函数将字符串s尾部按照函数参数f定义剪去后返回；
7. TrimSpace函数将字符串s的头部和尾部的空格或不显示字符全部剪去后返回；
8. TrimSuffix函数将字符换s的后缀字符串剪去后返回，注意后缀要整体匹配，而不是逐个字符匹配后剪切；
9. TrimPrefix函数将字符换s的前缀字符串剪去后返回，注意前缀要整体匹配，而不是逐个字符匹配后剪切；
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	println()

	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	println()
	println()
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))

	println()
	println()
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	println()
	println()
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Print(s)

	println()
	println()
	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))

	println()
	println()
	fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	println()
	println()
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))

	println()
	println()
	s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Print(s)

	println()

}
