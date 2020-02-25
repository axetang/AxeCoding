/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Index,strings.IndexAny,strings.IndexByte,strings.IndexFunc,
strings.IndexRune
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Index(s, substr string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
------------------------------------------------------------------------------------
**Description:
Index returns the index of the first instance of substr in s, or -1 if substr is not
present in s.
IndexAny returns the index of the first instance of any Unicode code point from chars
in s, or -1 if no Unicode code point from chars is present in s.
IndexByte returns the index of the first instance of c in s, or -1 if c is not present
in s.
IndexFunc returns the index into s of the first Unicode code point satisfying f(c),
or -1 if none do.
IndexRune returns the index of the first instance of the Unicode code point r, or -1
if rune is not present in s. If r is utf8.RuneError, it returns the first instance of
any invalid UTF-8 byte sequence.
-------------------------------------------------------------------------------------
**要点总结:
1. Index方法在s中查找substr第一次出现的位置。注意：位置是从0开始的；
2. IndexAny在s中查找substr中任意一个Unicode字符第一次出现的位置；
3. IndexByte方法在s中查找c byte第一次出现的位置；
4. IndexFunc方法在s中查找满足函数参数f(c)返回true的字符的位置；
5. IndexRune方法在s中查找r rune第一次出现的位置。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))

	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))

	fmt.Println(strings.IndexRune("chicke糖n", '糖'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
}
