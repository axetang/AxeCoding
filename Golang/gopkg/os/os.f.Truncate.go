/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Truncate
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Truncate(name string, size int64) error
 ------------------------------------------------------------------------------------
 **Description:
 Truncate changes the size of the named file. If the file is a symbolic link, it
 changes the size of the link's target. If there is an error, it will be of type
 *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Truncate函数改变文件name string的大小到参数size指定的大小；
 2. 如果文件是一个符号链接，则改变符号链接目标文件的大小。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The Hello.go's size is: %d\n", fi.Size())

	if err = os.Truncate("./_iofiles/Hello.go", 10); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Now the Hello.go's size is: %d\n", fi.Size())
}
