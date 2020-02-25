/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.Error
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Error struct {
    // Name is the file name for which the error occurred.
    Name string
    // Err is the underlying error.
    Err error
 }
 func (e *Error) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 Error is returned by LookPath when it fails to classify a file as an executable.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Error结构体用于LookPath函数返回错误，即无法找到可执行文件；
 2. Error方法实现了error接口，返回error字符串。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	e := &exec.Error{
		Name: "mv",
		Err:  errors.New("无法获取\"Hello\" 的文件状态(stat): 没有那个文件或目录"),
	}
	//test in ArchLinux
	fmt.Printf("%s\n", e.Error()) //exec: "mv": 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
}
