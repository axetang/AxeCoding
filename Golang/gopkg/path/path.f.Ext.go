/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Ext
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Ext(path string) string
 ------------------------------------------------------------------------------------
 **Description:
 Ext returns the file name extension used by path. The extension is the suffix
 beginning at the final dot in the final slash-separated element of path; it
 is empty if there is no dot.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Ext函数返回参数path中文件的扩展名；
 2. 如果没有扩展名，返回空串。
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
	fmt.Println(path.Ext("/a/b/c/bar"))     // ""
}
