/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.SyscallError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type SyscallError struct {
     Syscall string
     Err     error
 }
 func (e *SyscallError) Error() string
 func (e *SyscallError) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 SyscallError records an error from a specific system call.
 Timeout reports whether this error represents a timeout.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SyscallError记录一个指定系统调用发生的错误；
 2. Error方法获取SyscallError的字符串形式表达；
 3. TimeOut方法判断这个错误是否是Timeout错误。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	se := &os.SyscallError{
		Syscall: "/xx",
		Err:     errors.New("Syscall Error!"),
	}
	fmt.Printf("%s\n", se.Error())
	fmt.Printf("Timeout error? : \n", se.Timeout())
}
