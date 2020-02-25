/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.NewSyscallError
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewSyscallError(syscall string, err error) error
 ------------------------------------------------------------------------------------
 **Description:
 NewSyscallError returns, as an error, a new SyscallError with the given system call
 name and error details. As a convenience, if err is nil, NewSyscallError returns nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 这个函数创建一个系统错误；
 2. NewSyscallError方法返回一个错误，这个错误包含给定的system call的name和error detail；
 3. 如果err是nil，则函数返回nil；
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s\n", os.NewSyscallError("custom", errors.New("something wrong")))
	fmt.Printf("%s\n", os.NewSyscallError("custom", nil))
	fmt.Printf("%s\n", os.NewSyscallError("custom", errors.New("")))
}
