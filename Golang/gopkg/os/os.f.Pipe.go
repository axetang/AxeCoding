/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Pipe
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Pipe() (r *File, w *File, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Pipe returns a connected pair of Files; reads from r return bytes written to w. It
 returns the files and an error, if any.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Pipe函数返回两个*File，一个用作读，一个用作写，构成一个同步的管道。
 2. 该函数具体用途需后续分析后更新。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	r, w, _ := os.Pipe()
	fmt.Printf("r %T,%v;\nw %T,%v\n", r, r, w, w)
}
