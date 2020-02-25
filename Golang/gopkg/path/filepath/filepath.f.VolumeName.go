/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.VolumeName
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func VolumeName(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 VolumeName returns leading volume name. Given "C:\foo\bar" it returns "C:"
 on Windows. Given "\\host\share\foo" it returns "\\host\share". On other
 platforms it returns "".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. VolumeName函数用于Windows平台使用；
 2. 其他平台返回空串“”。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s := filepath.VolumeName("C:\foo\bar")
	fmt.Println(s)
}
