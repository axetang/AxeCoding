/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewWriter
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewWriter(w io.Writer) *Writer
 ------------------------------------------------------------------------------------
 **Description:
 NewWriter returns a new Writer whose buffer has the default size.
 ------------------------------------------------------------------------------------
 **要点总结:
 NewWriter返回一个新的，有默认Size缓存的Writer。
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	w.Write([]byte("hello,"))
	w.Write([]byte("world!"))
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}
