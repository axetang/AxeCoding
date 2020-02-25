/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.HTMLEscape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func HTMLEscape(w io.Writer, b []byte)
 ------------------------------------------------------------------------------------
 **Description:
 HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. HTMLEscape用于将参数b []byte转义后的HTML文本写入到w io.Writer中。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func main() {
	//NewBufferString返回一个*Buffer结构体，该结构体实现了io.Writer接口
	w := bytes.NewBufferString("")

	b := []byte("<script>alert('xss!')</script>")

	template.HTMLEscape(w, b)

	fmt.Println(w)

	template.HTMLEscape(os.Stdout, b)
	fmt.Println()

}
