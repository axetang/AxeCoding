/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.URLQueryEscape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func URLQueryEscaper(args ...interface{}) string
 ------------------------------------------------------------------------------------
 **Description:
 URLQueryEscaper returns the escaped value of the textual representation of its
 arguments in a form suitable for embedding in a URL query.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. URLQueryEscaper函数将变参args ...interface{}连接在一起后转义并返回转义后字符串；
 2. URL的转义规则又和HTML/JS的转移规则不同，需要了解后更新。
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

	fmt.Println(template.URLQueryEscaper(a, b, c))

}
