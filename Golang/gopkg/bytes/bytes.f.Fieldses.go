/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Fields,bytes.FieldsFunc
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Fields(s []byte) [][]byte
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
------------------------------------------------------------------------------------
**Description:
Fields interprets s as a sequence of UTF-8-encoded code points. It splits the slice s
around each instance of one or more consecutive white space characters, as defined by
unicode.IsSpace, returning a slice of subslices of s or an empty slice if s contains
only white space.
FieldsFunc interprets s as a sequence of UTF-8-encoded code points. It splits the
slice s at each run of code points c satisfying f(c) and returns a slice of subslices
of s. If all code points in s satisfy f(c), or len(s) == 0, an empty slice is returned.
FieldsFunc makes no guarantees about the order in which it calls f(c). If f does not
return consistent results for a given c, FieldsFunc may crash.
------------------------------------------------------------------------------------
**要点总结:
1. Fields把字节切片s按照一个或连续多个空白字符分隔成多个字节切片，如s只包含空白字符则返回空字节切片;
2. FieldsFunc把s解释为UTF-8编码的字符序列，对于每个Unicode字符c，如果函数参数f(c)返回true就把c
作为分隔字符对s进行拆分。如果所有都字符满足f(c)为true，则返回空的切片。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

//TestFieldsFunc eee
func TestFieldsFunc() {
	fields1 := bytes.FieldsFunc([]byte("ab,cd,,,e"), func(c rune) bool {
		return c == ','
	})
	fields2 := bytes.FieldsFunc([]byte("你好啊世界"), func(c rune) bool {
		return c == '啊'
	})
	fields3 := bytes.FieldsFunc([]byte(",,,"), func(c rune) bool {
		return c == ','
	})
	for i, s := range fields1 {
		fmt.Printf("%d: %s\n", i, string(s))
	}
	for i, s := range fields2 {
		fmt.Printf("%d: %s\n", i, string(s))
	}
	println("fileds3")
	for i, s := range fields3 {
		fmt.Printf("%d: %s\n", i, string(s))
	}

}

//TestFields dd
func TestFields() {
	s := []byte("a b\tc\rd\ne   f")
	for i, f := range bytes.Fields(s) {
		fmt.Printf("[%d]%s, %d\n", i, string(f), len(f))
	}
}

func main() {
	fmt.Println("--------------Fields----------------")
	TestFields()
	fmt.Println("--------------FieldsFunc----------------")
	TestFieldsFunc()
}
