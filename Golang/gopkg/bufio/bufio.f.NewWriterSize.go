/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewWriterSize
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewWriterSize(w io.Writer, size int) *Writer
 **Description:
 NewWriterSize returns a new Writer whose buffer has at least the specified size.
 If the argument io.Writer is already a Writer with large enough size, it returns
 the underlying Writer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewWriterSize返回一个新的Writer，它的缓存一定大于指定的size参数;
 2. 如果io.Writer参数已经是足够大的有缓存的Writer，函数就会返回它底层的Writer。
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	w.Write([]byte("hello,"))
	w.Write([]byte("world!"))
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}
