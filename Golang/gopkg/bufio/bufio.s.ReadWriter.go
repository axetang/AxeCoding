/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.s.ReadWriter
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ReadWriter struct {
    *Reader
    *Writer
 }
 ------------------------------------------------------------------------------------
 **Description:
 ReadWriter stores pointers to a Reader and a Writer. It implements io.ReadWriter.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 结构体ReadWriter存储输入输出指针。 它实现了io.ReadWriter；结构体ReadWriter的方法
 NewReadWriter分配新的ReadWriter来进行r和w的调度。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a new error!")
	if err != nil {
		fmt.Print(err)
	}
}
