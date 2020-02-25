/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.IsGraphic
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsGraphic(r rune) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsGraphic reports whether the rune is defined as a Graphic by Unicode. Such
 characters include letters, marks, numbers, punctuation, symbols, and spaces,
 from categories L, M, N, P, S, and Zs.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IsGraphic函数一个字符 rune r 是否是 unicode 图形字符。图形字符包括字母、标记、数字、符号、
 标点、空白；
 2. 此函数功能要在弄懂go语言unicode字符集后再更新。
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
