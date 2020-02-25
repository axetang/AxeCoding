/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Base
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Base(path string) string
 ------------------------------------------------------------------------------------
 **Description:
 Base returns the last element of path. Trailing slashes are removed before
 extracting the last element. If the path is empty, Base returns ".". If the
 path consists entirely of slashes, Base returns "/".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Base函数用来返回path的最后一个元素;
 2. 如果路径为空返回"."；
 3. 如果路径由斜线组成,返回"/"。
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))       // b
	fmt.Println(path.Base(""))           // .
	fmt.Println(path.Base("////"))       // /
	fmt.Println(path.Base("/a/"))        // b
	fmt.Println(path.Base("/Users/axe")) // b
	fmt.Println(path.Base("\a"))         // b
}
