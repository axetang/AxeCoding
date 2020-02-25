/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Abs
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Abs(path string) (string, error)
 -----------------------------------------------------------------------------------
 **Description:
 Abs returns an absolute representation of path. If the path is not absolute
 it will be joined with the current working directory to turn it into an
 absolute path. The absolute path name for a given file is not guaranteed
 to be unique. Abs calls Clean on the result.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Abs函数返回所给路径path的绝对路径string;
 2. Abs函数会对返回的路径结果调用path.Clean函数清理。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 当前路径为/home, 如下返回的path将会是/home/abs_demo.go
	path, _ := filepath.Abs("abs_demo.go")
	fmt.Printf("%v\n", path)
	path, _ = filepath.Abs("./_iofiles/README.md")
	fmt.Printf("%v\n", path)
	path, _ = filepath.Abs("../filepath/README.md")
	fmt.Printf("%v\n", path)
}
