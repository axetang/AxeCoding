/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Split
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Split(path string) (dir, file string)
 -----------------------------------------------------------------------------------
 **Description:
 Split splits path immediately following the final Separator, separating it
 into a directory and file name component. If there is no Separator in path,
 Split returns an empty dir and file set to path. The returned values have
 the property that path = dir+file.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Split函数和path.Split功能和使用方法相同，参看package path。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Split("static/myfile.css"))  // static/ myfile.css
	fmt.Println(filepath.Split("static/myfile.css/")) // static/ myfile.css
	fmt.Println(filepath.Split("static"))             // "" static
}
