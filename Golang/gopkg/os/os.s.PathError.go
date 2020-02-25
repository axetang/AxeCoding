/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.PathError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type PathError struct {
     Op   string
     Path string
     Err  error
 }
 func (e *PathError) Error() string
 func (e *PathError) Timeout() bool
------------------------------------------------------------------------------------
 **Description:
 PathError records an error and the operation and file path that caused it.
 Timeout reports whether this error represents a timeout.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. PathError记录了路径相关操作的错误信息；
 2. Error方法将PathError结构体实例的三个字段连接起来构成一个字符串，作为错误信息返回；
 3. Timeout方法判断是否这个错误是一个timeout错误。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	pe := &os.PathError{
		Op:   "cp",
		Path: "/oh/my/god",
		Err:  errors.New("path does'n exists!"),
	}
	fmt.Printf("%s\n", pe.Error())
	fmt.Println(pe.Timeout())
}
