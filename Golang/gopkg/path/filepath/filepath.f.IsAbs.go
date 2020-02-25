/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.IsAbs
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsAbs(path string) bool
 -----------------------------------------------------------------------------------
 **Description:
 IsAbs reports whether the path is absolute.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IsAbs函数和path.IsAbs函数功能用法一致，参见package path。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 当前路径为/home, 如下返回的path将会是/home/abs_demo.go
	path := filepath.IsAbs("abs_demo.go")
	fmt.Printf("%v\n", path)
}
