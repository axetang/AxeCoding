/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fprint
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fprint(w io.Writer, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Fprint formats using the default formats for its operands and writes to w. Spaces
 are added between operands when neither is a string. It returns the number of bytes
 written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fprint定义和实现如下,Fprintln把参数a切片中的值按照缺省格式写入io.Writer
 接口实例，操作符之间自动添加空格，但末尾不增加换行，最后返回写入的字节数
 func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a)
	n, err = w.Write(p.buf)
	p.free()
	return
 }
 2. Fprint调用了fmt包的私有函数newPrinter和doPrint，具体可查阅标准包源代码；
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	str := "abc"
	f := 3.14
	c := 1 + 2i
	t := true

	fmt.Fprint(os.Stdout, i, str, t)
	println()

	fmt.Fprint(os.Stdout, i, f, t)
	fmt.Println()

	fmt.Fprint(os.Stdout, i, f, c)
	fmt.Println()
}
