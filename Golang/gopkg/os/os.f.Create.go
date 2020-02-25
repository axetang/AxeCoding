/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Create
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Create(name string) (*File, error)
 ------------------------------------------------------------------------------------
 **Description:
 Create creates the named file with mode 0666 (before umask), truncating it if it
 already exists. If successful, methods on the returned File can be used for I/O; the
 associated file descriptor has mode O_RDWR. If there is an error, it will be of type
 *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Create方法创建一个文件，此方法会覆盖已存在的文件;
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Create("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("The %s has been created!\n", fi.Name())
}
