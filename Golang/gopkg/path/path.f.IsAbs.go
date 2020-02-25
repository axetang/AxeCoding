/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.IsAbs
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsAbs(path string) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsAbs reports whether the path is absolute.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IsAbs函数判断参数path是否是一个绝对路径。
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {

	fmt.Println(path.IsAbs("/home/zzz/go.pdf")) // true
	fmt.Println(path.IsAbs("home/zzz/go.pdf"))  // false
}
