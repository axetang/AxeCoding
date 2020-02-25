/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.ContainsRune
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ContainsRune(b []byte, r rune) bool
------------------------------------------------------------------------------------
**Description:
ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.
------------------------------------------------------------------------------------
**要点总结:
1. ContainsRune在b切片中搜索指定的rune字符r；
2. 注意，rune字符需要用单引号来标识，如果是双引号，代表的是字符串而不是字符。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("123飞45678")
	r1 := '飞'
	r2 := '哥'
	fmt.Println(bytes.ContainsRune(b, r1))
	fmt.Println(bytes.ContainsRune(b, r2))
}
