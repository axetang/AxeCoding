/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Join,bytes.Map
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Join(s [][]byte, sep []byte) []byte
func Map(mapping func(r rune) rune, s []byte) []byte
------------------------------------------------------------------------------------
**Description:
Join concatenates the elements of s to create a new byte slice. The separator sep is placed between elements in the resulting slice.
Map returns a copy of the byte slice s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement. The characters in s and the output are interpreted as UTF-8-encoded Unicode code points.
------------------------------------------------------------------------------------
**要点总结:
1. Join用sep把a中的每个字节切片连接成一个字节切片并返回。
2. Map把s解释为UTF-8编码的字节序列，并把s中的每个Unicode字符用mapping函数获得对应的字符并存放到
一个新创建的字节切片的对应位置，并返回此新创建的字节切片。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func TestJoin() {
	a := [][]byte{
		[]byte("hello"),
		[]byte("world"),
	}
	fmt.Println(string(bytes.Join(a, []byte(",*"))))
	//为什么[]byte(),有时候[]byte{}
	fmt.Println(string(bytes.Join(a, []byte{})))
	fmt.Println(string(bytes.Join(a, nil)))
}

func TestMap() {
	s1 := []byte("大家上午好")
	s2 := []byte("12345678")
	m1 := func(r rune) rune {
		if r == '上' {
			return '下'
		}
		return r
	}
	m2 := func(r rune) rune {
		return r + 1
	}
	fmt.Println(string(bytes.Map(m1, s1)))
	fmt.Println(string(s1))
	fmt.Println(string(bytes.Map(m2, s2)))
	fmt.Println(string(s2))
}

func main() {
	fmt.Println("------------Join-----------")
	TestJoin()
	fmt.Println("------------Map-----------")
	TestMap()
}
