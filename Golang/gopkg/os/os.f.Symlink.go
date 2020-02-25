/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Symlink
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Symlink(oldname, newname string) error
 ------------------------------------------------------------------------------------
 **Description:
 Symlink creates newname as a symbolic link to oldname. If there is an error, it will
 be of type *LinkError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Symlink函数给一个文件创建一个软链接；这和Link函数不同，Link创建硬链接；
 2. 返回错误*LinkError。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Symlink("./_iofiles/Hell.go", "./_iofiles/Lhell.go"); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("The Hell.go's  symbolic link Lhell.go has been created!\n")
}
