/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.Valid
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Valid(p []byte) bool
 ------------------------------------------------------------------------------------
 **Description:
 Valid reports whether p consists entirely of valid UTF-8-encoded runes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Valid方法判断参数p []byte是否组成了合法的UTF-8编码的runes。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}
