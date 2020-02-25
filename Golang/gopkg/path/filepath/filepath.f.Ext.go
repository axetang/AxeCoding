/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Ext
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Ext(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 Ext returns the file name extension used by path. The extension is the suffix
 beginning at the final dot in the final element of path; it is empty if there
 is no dot.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Ext函数和path.Ext函数用法和功能相同，参见package path。
*************************************************************************************/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
}
