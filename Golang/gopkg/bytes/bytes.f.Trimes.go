/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Trim,bytes.TrimFunc,bytes.TrimLeft,bytes.TrimLeftFunc,bytes.TrimPrefix,
bytes.TrimRight,bytes.TrimRightFunc,bytes.TrimSpace,bytes.TrimSuffix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Trim(s []byte, cutset string) []byte
func TrimFunc(s []byte, f func(r rune) bool) []byte
func TrimLeft(s []byte, cutset string) []byte
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func TrimRight(s []byte, cutset string) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
func TrimSpace(s []byte) []byte
func TrimPrefix(s, prefix []byte) []byte
func TrimSuffix(s, suffix []byte) []byte
------------------------------------------------------------------------------------
**Description:
Trim returns a subslice of s by slicing off all leading and trailing UTF-8-encoded code
points contained in cutset.
TrimFunc returns a subslice of s by slicing off all leading and trailing UTF-8-encoded
code points c that satisfy f(c).
TrimLeft returns a subslice of s by slicing off all leading UTF-8-encoded code points
contained in cutset.
TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing
off all leading UTF-8-encoded code points c that satisfy f(c).
TrimRight returns a subslice of s by slicing off all trailing UTF-8-encoded code
points that are contained in cutset.
TrimRightFunc returns a subslice of s by slicing off all trailing UTF-8-encoded code
points c that satisfy f(c).
TrimSpace returns a subslice of s by slicing off all leading and trailing white space,
as defined by Unicode.
TrimPrefix returns s without the provided leading prefix string. If s doesn't start
with prefix, s is returned unchanged.
TrimSuffix returns s without the provided trailing suffix string. If s doesn't end
with suffix, s is returned unchanged.
------------------------------------------------------------------------------------
**要点总结:
1. Trim返回s的子字节切片，s的头部和尾部被剪切到都不包含cutset字符串中的任意字符；
2. TrimFunc功能和Trim类似，但通过函数参数f来判断f(r)是否为true的方式，剪切s的头部和尾部；
3. TrimLeft返回s的子字节切片，s的头部被剪切到都不包含cutset字符串中的任意字符；
4. TrimLeftFunc功能和TrimLeft类似，但通过函数参数f来判断f(r)是否为true的方式，剪切s的头部；
5. TrimRight返回s的子字节切片，s的尾部被剪切到都不包含cutset字符串中的任意字符；
6. TrimRightFunc功能和TrimLeft类似，但通过函数参数f来判断f(r)是否为true的方式，剪切s的尾部；
7. TrimSpace返回s的子字节切片，s的头部和尾部被剪切到不包含任何空格类字符；
8. TrimPrefix返回s的子字节切片，s的前缀如果包含prefix则被剪掉后返回，注意是要整体匹配而不是匹配任意字节；
9. TrimSuffix返回s的子字节切片，s的后缀如果包含suffix则被剪掉后返回，注意是要整体匹配而不是匹配任意字节。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("-----Trim-----")
	s := []byte("11234567898")
	fmt.Println(string(bytes.Trim(s, "1389")))

	fmt.Println("-----TrimFunc-----")
	s = []byte("12345678")
	f := func(r rune) bool {
		return r <= '3' || r >= '6'
	}
	fmt.Println(string(bytes.TrimFunc(s, f)))

	fmt.Println("-----TrimLeft-----")
	s = []byte("9812345")
	fmt.Println(string(bytes.TrimLeft(s, "129")))

	fmt.Println("-----TrimLeftFunc-----")
	s = []byte("123456789")
	f = func(r rune) bool {
		return r < '4'
	}
	fmt.Println(string(bytes.TrimLeftFunc(s, f)))

	fmt.Println("-----TrimRight-----")
	s = []byte("98123456789")
	fmt.Println(string(bytes.TrimRight(s, "689")))

	fmt.Println("-----TrimRightFunc-----")
	s = []byte("123456789")
	f = func(r rune) bool {
		return r > '7'
	}
	fmt.Println(string(bytes.TrimRightFunc(s, f)))

	fmt.Println("-----TrimSpace-----")
	s1 := bytes.TrimSpace([]byte("   \t  hello, world!\r\n"))
	s2 := bytes.TrimSpace([]byte("   \t  \r\n"))
	fmt.Printf("%d %s\n", len(s1), string(s1))
	fmt.Printf("%d %s\n", len(s2), string(s2))

	fmt.Println("-----TrimPrefix-----")
	s1 = bytes.TrimPrefix([]byte("prefix hello, world!\r\n"), []byte("prefix"))
	s2 = bytes.TrimPrefix([]byte("pprefix helo, world\r\n"), []byte("prefix"))
	fmt.Printf("%d %s\n", len(s1), string(s1))
	fmt.Printf("%d %s\n", len(s2), string(s2))

	fmt.Println("-----TrimSuffix-----")
	s1 = bytes.TrimSuffix([]byte("prefix hello, world!"), []byte("world!"))
	s2 = bytes.TrimSuffix([]byte("pprefix helo, world\r\n"), []byte("world!"))
	fmt.Printf("%d %s\n", len(s1), string(s1))
	fmt.Printf("%d %s\n", len(s2), string(s2))
}
