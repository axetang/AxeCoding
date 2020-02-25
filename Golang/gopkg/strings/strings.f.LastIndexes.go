/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.LastIndex,strings.LastIndexAny,strings.LastIndexByte,
strings.LastIndexFunc
**Type: func
------------------------------------------------------------------------------------
**Definition:
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
------------------------------------------------------------------------------------
**Description:
LastIndex returns the index of the last instance of substr in s, or -1 if substr is
not present in s.
LastIndexAny returns the index of the last instance of any Unicode code point from
chars in s, or -1 if no Unicode code point from chars is present in s.
LastIndexByte returns the index of the last instance of c in s, or -1 if c is not
present in s.
LastIndexFunc returns the index into s of the last Unicode code point satisfying
f(c), or -1 if none do.
------------------------------------------------------------------------------------
**要点总结:
1. LastIndex方法返回substr在s中最后一次出现的位置；
2. LastIndexAny方法返回substr中任意一个Unicode字符在s中最后一次出现的位置；
3. LastIndexByte方法返回c byte在s中最后一次出现的位置；
4. LastIndexFunc方法和LastIndex方法类似，但通过函数参数f来判断满足f(c)最后出现的位置.
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))

	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))

	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))

	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))

}
