/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.HTMLEscaper
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func HTMLEscaper(args ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 HTMLEscaper returns the escaped HTML equivalent of the textual representation of
 its arguments.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. HTMLEscaper返回多个变参args ...interface{}一起转义后的字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"html/template"
)

func main() {

	a := "<script>alert('xss!')</script>"

	b := "<h1>Raymond</h1>"

	c := "<h2>Chou</h2>"

	fmt.Println(template.HTMLEscaper(a, b, c))

}
