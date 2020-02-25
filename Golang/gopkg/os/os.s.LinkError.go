/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.LinkError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type LinkError struct {
     Op  string
     Old string
     New string
     Err error
 }
 func (e *LinkError) Error() string
------------------------------------------------------------------------------------
 **Description:
 LinkError records an error during a link or symlink or rename system call and the
 paths that caused it.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LinkError记录link，symlink或rename等系统调用时的错误；
 2. Error方法实现了系统接口error，这个函数将LinkError结构体实例的四个参数连接起来构成一个字符串，
 并作为错误信息返回。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	le := &os.LinkError{
		Op:  "ln",
		Old: "old",
		New: "new",
		Err: errors.New("ln Error!"),
	}
	fmt.Printf("%s\n", le.Error())
}
