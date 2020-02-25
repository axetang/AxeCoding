/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.NewFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewFile(fd uintptr, name string) *File
 ------------------------------------------------------------------------------------
 **Description:
 NewFile returns a new File with the given file descriptor and name. The returned
 value will be nil if fd is not a valid file descriptor. On Unix systems, if the file
 descriptor is in non-blocking mode, NewFile will attempt to return a pollable File
 (one for which the SetDeadline methods work).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewFile方法新建一个文件（但不保存），返回文件指针；
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("type：%T, Value: %v\n", fi.Fd(), fi.Fd())
	file := os.NewFile(fi.Fd(), "./_iofiles/World.go")
	defer file.Close()
	fmt.Printf("The %s has been new!\n", file.Name())
}
