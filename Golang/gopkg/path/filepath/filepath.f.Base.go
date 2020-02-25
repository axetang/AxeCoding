/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Base
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Base(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 Base returns the last element of path. Trailing path separators are removed
 before extracting the last element. If the path is empty, Base returns ".".
 If the path consists entirely of separators, Base returns a single separator.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 和path.Base功能相同；
 2. Base函数用来返回path的最后一个元素;
 3. 如果路径为空返回"."；
 4. 如果路径由斜线组成,返回"/"
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("/a/b"))       // b
	fmt.Println(filepath.Base(""))           // .
	fmt.Println(filepath.Base("////"))       // /
	fmt.Println(filepath.Base("/a/"))        // b
	fmt.Println(filepath.Base("/Users/axe")) // b
	fmt.Println(filepath.Base("\a"))         // b
}
