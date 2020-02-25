/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Open
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Open(name string) (*File, error)
 ------------------------------------------------------------------------------------
 **Description:
 opens the named file with specified flag (O_RDONLY etc.) and perm (before umask), if
 applicable. If successful, methods on the returned File can be used for I/O. If there
 is an error, it will be of type *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Open方法以只读的模式打开一个文件；
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
	data := make([]byte, 100)
	fi.Read(data)
	fmt.Printf("The %s's content is: \n%s \n", fi.Name(), data)
}
