/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.EvalSymlinks
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func EvalSymlinks(path string) (string, error)
 -----------------------------------------------------------------------------------
 **Description:
 EvalSymlinks returns the path name after the evaluation of any symbolic links.
 If path is relative the result will be relative to the current directory,
 unless one of the components is an absolute symbolic link. EvalSymlinks calls
 Clean on the result.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. EvalSymlinks函数返回一个链接文件的实际的路径；
 2. 例如在/home下创建一个link.log的文件，然后
 	cd xuchdong;
	ln -s /home/link.log link，
	使用该函数可以找到原始的文件即/home/link.log
 3. 下面例子程序执行结果和预计不符，还需再看看。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 环境准备：首先在/home目录下创建一个link.log的文件,
	// 然后在当前目录下使用ln -s /home/link.log link_other

	path, _ := filepath.EvalSymlinks("link_other")
	fmt.Printf("path is %s\n", path) // /home/link.log
}
