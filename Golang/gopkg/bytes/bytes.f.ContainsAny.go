/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.ContainsAny
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ContainsAny(b []byte, chars string) bool
------------------------------------------------------------------------------------
**Description:
ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.
------------------------------------------------------------------------------------
**要点总结:
1. ContainsAny函数从字节切片b中查找字符串chars中的任意一个字符，如果存在返回true，否则返回false
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("12345678")
	s1 := "456"
	s2 := "901"
	fmt.Println(bytes.ContainsAny(b, s1))
	fmt.Println(bytes.ContainsAny(b, s2))
}
