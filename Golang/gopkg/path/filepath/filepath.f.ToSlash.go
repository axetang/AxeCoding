/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.ToSlash
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ToSlash(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 ToSlash returns the result of replacing each separator character in path with
 a slash ('/') character. Multiple separators are replaced by multiple slashes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ToSlash函数将所有的路径分隔符使用"/"替换；
 2. 分隔符见os.PathSeparator, linux默认的分隔符是"/";
 3. 这个函数具体用途要继续了解下后更新。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.ToSlash("staticmy///:file.css"))
}
