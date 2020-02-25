/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Rel
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Rel(basepath, targpath string) (string, error)
 -----------------------------------------------------------------------------------
 **Description:
 Rel returns a relative path that is lexically equivalent to targpath when
 joined to basepath with an intervening separator. That is, Join(basepath,
 Rel(basepath, targpath)) is equivalent to targpath itself. On success, the
 returned path will always be relative to basepath, even if basepath and
 targpath share no elements. An error is returned if targpath can't be made
 relative to basepath or if knowing the current working directory would be
 necessary to compute it. Rel calls Clean on the result.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Rel函数返回以basepath为基准的相对路径。
 2. Join(basepath, Rel(basepath, targpath)) 等于targpath;
 3. 这个函数返回值需要仔细评估清楚后使用，避免出现意外结果。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 返回 xuchdong/rel_demo.go
	path, _ := filepath.Rel("/home", "/home/xuchdong/rel_demo.go")
	fmt.Printf("%s\n", path)
	path, _ = filepath.Rel("/Users", "/home/xuchdong/rel_demo.go")
	fmt.Printf("%s\n", path)
	path, _ = filepath.Rel("/xuchdong", "/home/xuchdong/rel_demo.go")
	fmt.Printf("%s\n", path)
}
