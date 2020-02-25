/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Dir
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Dir(path string) string
 ------------------------------------------------------------------------------------
 **Description:
 Dir returns all but the last element of path, typically the path's directory.
 After dropping the final element using Split, the path is Cleaned and trailing
 slashes are removed. If the path is empty, Dir returns ".". If the path
 consists entirely of slashes followed by non-slash bytes, Dir returns a
 single slash. In any other case, the returned path does not end in a slash.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Dir函数返回path中最后一个元素的路径
 2. Dir函数规则：
		1）通常是路径最后一个元素的路径目录
		2）路径为空返回.
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/a/b/c"))  // /a/b
	fmt.Println(path.Dir("/a/b/c/")) // /a/b/c
	fmt.Println(path.Dir(""))        // .
}
