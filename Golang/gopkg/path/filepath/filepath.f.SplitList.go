/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.SplitList
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SplitList(path string) []string
 -----------------------------------------------------------------------------------
 **Description:
 SplitList splits a list of paths joined by the OS-specific ListSeparator,
 usually found in PATH or GOPATH environment variables. Unlike strings.Split,
 SplitList returns an empty slice when passed an empty string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 将路径字符串使用路径列表分隔符分开。路径列表分隔符见os.PathListSeparator, linux的
路径列表分隔符是":", windows的路径列表分隔符是";"，主要用于PATH或是GOPATH等环境变
量。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Printf("%v\n", filepath.SplitList("a/b:c/d")) //[a/b; c/d]
}
