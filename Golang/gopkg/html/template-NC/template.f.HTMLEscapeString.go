/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.HTMLEscapeString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func HTMLEscapeString(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. HTMLEscapeString返回参数s string转义后的HTML文本.
*************************************************************************************/
package main

import (
	"fmt"
	"html/template"
)

func main() {

	s := "<script>alert('xss!')</script>"

	fmt.Println(template.HTMLEscapeString(s))

}
