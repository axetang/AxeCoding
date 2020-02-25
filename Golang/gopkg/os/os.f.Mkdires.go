/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Mkdir,os.MkdirAll
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Mkdir(name string, perm FileMode) error
 func MkdirAll(path string, perm FileMode) error
 ------------------------------------------------------------------------------------
 **Description:
 Mkdir creates a new directory with the specified name and permission bits (before
 umask). If there is an error, it will be of type *PathError.
 MkdirAll creates a directory named path, along with any necessary parents, and
 returns nil, or else returns an error. The permission bits perm (before umask) are
 used for all directories that MkdirAll creates. If path is already a directory,
 MkdirAll does nothing and returns nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Mkdir函数依据指定的名称name和权限perm创建一个目录，如果出错返回*PathError；如果要创建的目录的
 父目录不存在，则返回错误；要解决这个问题，需要用到下面的函数MkdirAll；
 2. MkdirAll函数依据指定的路径创建所有目录，perm应用于所有目录。如果目录已经存在，函数不执行任何操作
 并返回nil。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("---Mkdir---")
	if err := os.Mkdir("./_iofiles/hellodir", 0777); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("./_iofiles/hellodir has been created!\n")
		fmt.Println("err is ", err)
	}
	fmt.Println("---Mkdir 2---")
	if err := os.Mkdir("./_files/hellodir", 0777); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("./_iofiles/hellodir has been created!\n")
		fmt.Println("err is ", err)
	}

	fmt.Println("---MkdirAll---")
	if err := os.MkdirAll("./_iofiles/hellodir/world", 0777); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("err is ", err)
		fmt.Printf("./_iofiles/hellodir/world has been created!")
	}
}
