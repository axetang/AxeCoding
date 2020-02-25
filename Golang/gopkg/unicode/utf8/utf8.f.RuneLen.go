/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.RuneLen
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func RuneLen(r rune) int
 ------------------------------------------------------------------------------------
 **Description:
 RuneLen returns the number of bytes required to encode the rune. It returns -1 if
 the rune is not a valid value to encode in UTF-8.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneLen方法返回参数r rune的字节数；
 2. 如果参数r是一个非法的UTF-8码值，则返回错-1.
*************************************************************************************/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
}
