/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Rename
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Rename(oldpath, newpath string) error
 ------------------------------------------------------------------------------------
 **Description:
 Rename renames (moves) oldpath to newpath. If newpath already exists and is not a
 directory, Rename replaces it. OS-specific restrictions may apply when oldpath and
 newpath are in different directories. If there is an error, it will be of type
 *LinkError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Rename函数将oldpath移动到newpath，即可以修改文件名并移动，也可以修改目录名并移动；
 2. 如果newpath已经存在且不是一个目录，Rename会覆盖它；
 3. 返回LinkError错误。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("./_iofiles/Hello.go", "./_iofiles/World.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Hello.go has been renamed to World.go!\n")
	err = os.Rename("./_iofiles/hellodir", "./_iofiles/Worlddir")
	fmt.Printf("Hellodir has been renamed to Worlddir!\n")
}
