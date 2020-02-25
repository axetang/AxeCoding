/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.JSEscape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func JSEscape(w io.Writer, b []byte)
 ------------------------------------------------------------------------------------
 **Description:
 JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. JSEscape将转义后的b []byte写入w io.Writer；
 2. 注意：JS的转义方法和HTML不同，后续查阅相关信息后更新。
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

	template.JSEscape(w, b)

	fmt.Println(w)

	template.JSEscape(os.Stdout, b)
	fmt.Println()

}
