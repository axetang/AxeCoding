/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fprintf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Fprintf formats according to a format specifier and writes to w. It returns the
 number of bytes written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fprintf定义和实现如下,Fprintf把参数a切片中的值按照指定格式写入io.Writer
 接口实例，最后返回写入的字节数
 func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
 }
 2. Fprintf调用了fmt包的私有函数newPrinter和doPrintf，具体可查阅标准包源代码；
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

	fmt.Fprintf(os.Stdout, "%d %s %6.5f %v %t %g %G %e %E", i, str, f, c, t, f, f, f, f)
	println()

}
