/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.JSEscapeString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func JSEscapeString(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 JSEscapeString returns the escaped JavaScript equivalent of the plain text
 data s.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. JSEscapeString将参数s string进行JS转义并返回转义后的字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"html/template"
)

func main() {

	s := "<script>alert('xss!')</script>"

	fmt.Println(template.JSEscapeString(s))

}
