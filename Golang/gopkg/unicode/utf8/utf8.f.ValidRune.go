/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.ValidRune
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ValidRune(r rune) bool
 ------------------------------------------------------------------------------------
 **Description:
 ValidRune reports whether r can be legally encoded as UTF-8. Code points that are
 out of range or a surrogate half are illegal.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ValidRune函数判断参数r rune字符是否是一个合法的UTF-8编码字符；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := 'a'
	invalid := rune(0xfffffff)

	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
}
