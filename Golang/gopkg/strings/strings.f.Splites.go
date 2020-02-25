/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Split,strings.SplitAfter,strings.SplitAfterN,strings.SplitN
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
------------------------------------------------------------------------------------
**Description:
Split slices s into all substrings separated by sep and returns a slice of the
substrings between those separators. If s does not contain sep and sep is not empty,
Split returns a slice of length 1 whose only element is s. If sep is empty, Split
splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty
slice. It is equivalent to SplitN with a count of -1.
SplitAfter slices s into all substrings after each instance of sep and returns a slice
of those substrings. If s does not contain sep and sep is not empty, SplitAfter returns a slice of length
1 whose only element is s. If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are
empty, SplitAfter returns an empty slice. It is equivalent to SplitAfterN with a count
of -1.
SplitAfterN slices s into substrings after each instance of sep and returns a slice
of those substrings.
The count determines the number of substrings to return:
	n > 0: at most n substrings; the last substring will be the unsplit remainder.
	n == 0: the result is nil (zero substrings)
	n < 0: all substrings
Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.
SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
The count determines the number of substrings to return:
	n > 0: at most n substrings; the last substring will be the unsplit remainder.
	n == 0: the result is nil (zero substrings)
	n < 0: all substrings
Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for Split.
-------------------------------------------------------------------------------------
**要点总结:
1. Split函数把s根据sep分割，返回分割之后子字符串的slice。如果s中没有sep，s也不为空，则返回s；
如果sep为空，那么每一个字符都要分割;
2. SplitAfter函数把s根据sep分割，返回分割之后子字符串的slice,和split一样只是返回的子字符串保留sep，
如果sep为空，那么每一个字符都分割;
3. SplitAfterN函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留
sep，如果sep为空，那么每一个字符都分割;参数n表示最多可以分割的字符串数量；
	- n > 0: 最多n个子字符串; 最后一个就是剩下未分割的子字符串.
	- n == 0: 返回为空字符串
	- n < 0: 返回所有的子字符串，和SplitAfter
4. SplitN函数s根据sep分割，返回分割之后子字符串的slice，返回的子串的长度如n的定义，如果sep为空，
那么每一个字符都分割。
	- n > 0: 最多n个子字符串; 最后一个就是剩下未分割的子字符串.
	- n == 0: 返回空字符串
	- n < 0: 返回所有的子字符串，和SplitAfter
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a,b,c", ";"))
	fmt.Printf("%q\n", strings.Split("", ""))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	println()
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) //["a," "b," "c"]

	println()
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) //["a," "b,c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", "", 5))  //["a" "," "b" "," "c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 1)) //["a,b,c"]

	println()
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) //["a" "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) //[] (nil = true)

}
