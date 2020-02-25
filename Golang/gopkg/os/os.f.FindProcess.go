/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.FindProcess
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FindProcess(pid int) (*Process, error)
 ------------------------------------------------------------------------------------
 **Description:
 FindProcess looks for a running process by its pid.
 The Process it returns can be used to obtain information about the underlying
 operating system process.
 On Unix systems, FindProcess always succeeds and returns a Process for the given pid,
 regardless of whether the process exists.
------------------------------------------------------------------------------------
 **要点总结:
 1. FindProcess函数根据参数pid int查找并返回进程对象*Process；
 2. Unix系统下，FindProcess函数总是执行成功并返回一个给定pid的进程，无论这个进程是否存在。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	p, err := os.FindProcess(31209)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process info: %+v\n", p)

	p, err = os.FindProcess(88888)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process info: %+v\n", p)

}
