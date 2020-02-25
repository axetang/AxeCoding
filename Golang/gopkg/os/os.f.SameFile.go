/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.SameFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SameFile(fi1, fi2 FileInfo) bool
 ------------------------------------------------------------------------------------
 **Description:
 SameFile reports whether fi1 and fi2 describe the same file. For example, on Unix
 this means that the device and inode fields of the two underlying structures are
 identical; on other systems the decision may be based on the path names. SameFile
 only applies to results returned by this package's Stat. It returns false in other
 cases.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SameFile函数判断两个文件信息fi1和fi2是否相同；

*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi1, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fi2, err := os.Stat("./_iofiles/World.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%t\n", os.SameFile(fi1, fi2))
	fmt.Printf("%t\n", os.SameFile(fi1, fi1))
}
