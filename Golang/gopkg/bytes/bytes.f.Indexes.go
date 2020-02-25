/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Index,bytes.IndexAny,bytes.IndexByte,bytes.IndexFunc,bytes.IndexRune
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Index(s, sep []byte) int
func IndexAny(s []byte, chars string) int
func IndexByte(s []byte, c byte) int
func IndexFunc(s []byte, f func(r rune) bool) int
func IndexRune(s []byte, r rune) int
------------------------------------------------------------------------------------
**Description:
Index returns the index of the first instance of sep in s, or -1 if sep is not present
in s.
IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points. It returns
the byte index of the first occurrence in s of any of the Unicode code points in chars.
It returns -1 if chars is empty or if there is no code point in common.
IndexByte returns the index of the first instance of c in s, or -1 if c is not present
in s.
IndexFunc interprets s as a sequence of UTF-8-encoded Unicode code points. It returns
the byte index in s of the first Unicode code point satisfying f(c), or -1 if none do.
IndexRune interprets s as a sequence of UTF-8-encoded Unicode code points. It returns
the byte index of the first occurrence in s of the given rune. It returns -1 if rune
is not present in s.
------------------------------------------------------------------------------------
**要点总结:
1. Index返回sep在s中第一次出现的位置索引，如果sep不在s中则返回-1;
2. IndexAny把s解释为UTF-8编码的字节序列，返回chars中任一个字符在s中第一次出现的位置索引；如果s中
不包含chars中任何一个字符则返回-1;
3. IndexByte检查c在s中第一次出现的位置索引；如果s中不包含c则返回-1;
4. IndexRune把s解释为UTF-8字节序列，并返回r在s中的位置索引；如果s中不包含r则返回-1;
5. IndexFunc把s解释成UTF-8字节序列，并返回第一个满足f(c)==true的字符c的位置索引；如果没有字符
满足f(c)==true则返回-1.
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func TestIndex() {
	s := []byte("1234,12345678")
	fmt.Println(bytes.Index(s, []byte("34")))
	fmt.Println(bytes.Index(s, []byte("789")))
}

func TestIndexAny() {
	s := []byte("大家好大家早")
	index := bytes.IndexAny(s, "大好早")
	if index >= 0 {
		fmt.Printf("%d: %s\n", index, string(s[index:]))
	}
}

func TestIndexByte() {
	fmt.Println(bytes.IndexByte([]byte("1234"), '3'))
}

func TestIndexRune() {
	s := []byte("你好世界")
	index := bytes.IndexRune(s, '好')
	if index >= 0 {
		fmt.Println(index)
		fmt.Println(string(s[index:]))
	}
}

func TestIndexFunc() {
	s := []byte("你好世界")
	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		return r == '好'
	}))
	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		return r == '!'
	}))
}

func main() {
	fmt.Println("-----------Index--------------")
	TestIndex()
	fmt.Println("-----------IndexAny--------------")
	TestIndexAny()
	fmt.Println("-----------IndexByte--------------")
	TestIndexByte()
	fmt.Println("-----------IndexRune--------------")
	TestIndexRune()
	fmt.Println("-----------IndexFunc--------------")
	TestIndexFunc()
}
