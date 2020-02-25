/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Print
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Print(a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Print formats using the default formats for its operands and writes to standard
 output. Spaces are added between operands when neither is a string. It returns the
 number of bytes written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Print使用其操作数的默认格式进行格式化并写入到标准输出;
 2. 当两个连续的操作数均不为字符串时，它们之间就会添加空格;
 3. 它返回写入的字节数以及任何遇到的错误。
 4. 查看标准包源码，Print定义和实现如下
 func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
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
		fmt.Print(err)
	}
	println()
}
