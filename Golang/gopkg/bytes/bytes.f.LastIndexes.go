/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.LastIndex,bytes.,bytes.,bytes.
**Type: func
------------------------------------------------------------------------------------
**Definition:
func LastIndex(s, sep []byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexByte(s []byte, c byte) int
func LastIndexFunc(s []byte, f func(r rune) bool) int
------------------------------------------------------------------------------------
**Description:
LastIndex returns the index of the last instance of sep in s, or -1 if sep is not
present in s.
LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code points. It
returns the byte index of the last occurrence in s of any of the Unicode code points
in chars. It returns -1 if chars is empty or if there is no code point in common.
LastIndexByte returns the index of the last instance of c in s, or -1 if c is not
present in s.
LastIndexFunc interprets s as a sequence of UTF-8-encoded Unicode code points. It
returns the byte index in s of the last Unicode code point satisfying f(c), or -1
if none do.
------------------------------------------------------------------------------------
**要点总结:
1. LastIndex返回sep在s中最后一次出现的位置索引，如果s中不包含sep则返回-1;
2. LastIndexAny把s解释为UTF-8编码的字节序列，返回chars中任意字符在s中最后出现的位置索引。如果
chars为空或者s中不包含chars中的任意字符，则返回-1;
3. LastIndexByte返回c在s中最后一次出现的位置索引，如果s中不包括c则返回-1；
4. LastIndexFunc把s解释为UTF-8编码的字节序列，返回满足函数参数f(c)返回true的字符c在s中最后一次
出现的位置索引, 如果没有找到这样的字符则返回-1。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func TestLastIndex() {
	s := []byte("abcd,abcd,cabcd")
	fmt.Println(bytes.LastIndex(s, []byte("abc")))
	fmt.Println(bytes.LastIndex(s, []byte(",c")))
}

func TestLastIndexAny() {
	s := []byte("123456789")
	fmt.Println(bytes.LastIndexAny(s, "789"))
	fmt.Println(bytes.LastIndexAny(s, "987"))
	fmt.Println(bytes.LastIndexAny(s, "0"))
	fmt.Println(bytes.LastIndexAny(s, ""))
	fmt.Println(bytes.LastIndexAny(s, ""))
}

func TestLastIndexByte() {
	s := []byte("abcd,abcd")
	fmt.Println(bytes.LastIndexByte(s, byte('c')))
}

func TestLastIndexFunc() {
	f1 := func(r rune) bool { return r == '7' }
	f2 := func(r rune) bool { return r == '9' }
	s := []byte("12345678")
	fmt.Println(bytes.LastIndexFunc(s, f1))
	fmt.Println(bytes.LastIndexFunc(s, f2))
}

func main() {
	fmt.Println("-------LastIndex--------")
	TestLastIndex()
	fmt.Println("-------LastIndexAny--------")
	TestLastIndexAny()
	fmt.Println("-------LastIndexByte--------")
	TestLastIndexByte()
	fmt.Println("-------LastIndexFunc--------")
	TestLastIndexFunc()
}
