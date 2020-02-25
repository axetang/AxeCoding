/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewReadWriter
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReadWriter(r *Reader, w *Writer) *ReadWriter
 ------------------------------------------------------------------------------------
 **Description:
 NewReadWriter allocates a new ReadWriter that dispatches to r and w.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 结构体ReadWriter存储输入输出指针。 它实现了io.ReadWriter；
 2. 结构体ReadWriter的方法NewReadWriter分配新的ReadWriter来进行r和w的调度。
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	wb := bytes.NewBuffer(nil)
	r := bufio.NewReader(rb)
	w := bytes.NewBuffer(wb)
	rw := bufio.NewReadWriter(r, w)
	// use rw to read
	var rbuf [128]byte
	if n, err := rw.Read(rbuf[:]); err != nil {
		return
	} else {
		fmt.Println(string(rbuf[:n]))
	}
	// use rw to write
	if _, err := rw.Write([]byte("a string to be written")); err != nil {
		return
	}
	rw.Flush()
	fmt.Println(string(wb.Bytes()))
}
