/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.TempDir
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func TempDir() string
 ------------------------------------------------------------------------------------
 **Description:
 TempDir returns the default directory to use for temporary files.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. TempDir函数返回系统的缺省临时文件目录。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The temp dir is: %s\n", os.TempDir())
}
