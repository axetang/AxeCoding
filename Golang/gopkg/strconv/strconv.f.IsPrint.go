/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.IsPrint
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsPrint(r rune) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsPrint reports whether the rune is defined as printable by Go, with the same
 definition as unicode.IsPrint: letters, numbers, punctuation, symbols and ASCII
 space.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IsPrint判断参数字符r rune是否是Go语言的可打印字符，如：字母、数字、标点符号和AscII空格。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)
}
