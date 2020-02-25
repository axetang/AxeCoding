/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Println
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Println(a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Println formats using the default formats for its operands and writes to standard
 output. Spaces are always added between operands and a newline is appended. It
 returns the number of bytes written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Println使用其操作数的默认格式进行格式化并写入到标准输出, 其操作数之间总是添加空格，且总在最后
 追加一个换行符, 它返回写入的字节数以及任何遇到的错误。
 2. 查看标准包源码，Println定义和实现如下
  func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
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
		fmt.Println(err)
	}
}
