/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Link
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Link(oldname, newname string) error
 ------------------------------------------------------------------------------------
 **Description:
 Link creates newname as a hard link to the oldname file. If there is an error, it
 will be of type *LinkError.
 ------------------------------------------------------------------------------------
 **要点总结:
 Link函数给一个文件oldname建立一个硬链接
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Link("./_iofiles/Hello.go", "./_iofiles/Lhello.go"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Lhello.go has created!\n")
}
