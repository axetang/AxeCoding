/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Printf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Printf(format string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Printf formats according to a format specifier and writes to standard output. It
 returns the number of bytes written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Printf 根据于格式说明符进行格式化并写入到标准输出,它返回写入的字节数以及任何遇到的写入错误;
 2. 查看标准包源码，Printf定义和实现如下
   func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
 }
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a new error!")
	if err != nil {
		fmt.Printf("%s", err)
	}
	println()
}
