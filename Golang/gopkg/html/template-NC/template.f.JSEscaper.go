/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.JSEscaper
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func JSEscaper(args ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 JSEscaper returns the escaped JavaScript equivalent of the textual representation
 of its arguments.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. JSEscaper将变参args ...interface{}连接后一起进行JS转义并返回结果字符串。
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

	fmt.Println(template.JSEscaper(a, b, c))

}
