/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Readlink
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Readlink(name string) (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 Readlink returns the destination of the named symbolic link. If there is an error,
 it will be of type *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 如下程序因何出错，需要再审视下；
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	link, err := os.Readlink("./_iofiles/Lhello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", link)
}
