/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Remove(name string) error
 func RemoveAll(path string) error
 ------------------------------------------------------------------------------------
 **Description:
 Remove removes the named file or (empty) directory. If there is an error, it will
 be of type *PathError.
 RemoveAll removes path and any children it contains. It removes everything it can
 but returns the first error it encounters. If the path does not exist, RemoveAll
 returns nil (no error).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Remove函数删除一个文件或空目录；
 2. RemoveAll函数删除一个目录及其所有子目录；
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("./_iofiles/hellodir/He.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("./_iofiles/hellodir/He.go has been removed!\n")

	err = os.RemoveAll("./_iofiles/hellodir/world")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("./_iofiles/hellodir/world has been removed!\n")

}
