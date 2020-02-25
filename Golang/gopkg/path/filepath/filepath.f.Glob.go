/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Glob
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Glob(pattern string) (matches []string, err error)
 -----------------------------------------------------------------------------------
 **Description:
 Glob returns the names of all files matching pattern or nil if there is no
 matching file. The syntax of patterns is the same as in Match. The pattern
 may describe hierarchical names such as /usr/ * /bin/ed (assuming the
 Separator is '/').
 Glob ignores file system errors such as I/O errors reading directories.
 The only possible returned error is ErrBadPattern, when pattern is malformed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Glob函数返回目录下所有匹配pattern string的文件；
 2. 这里的pattern用法和Match函数一致，可参看match函数的pattern约定。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//返回目录下所有的go文件
	matches, _ := filepath.Glob("*.go")
	fmt.Println(matches, "\n\n")

	//找出/home/ 目录下的所有的log文件
	matches, _ = filepath.Glob("/Users/axe/*.log")
	fmt.Println(matches)
}
