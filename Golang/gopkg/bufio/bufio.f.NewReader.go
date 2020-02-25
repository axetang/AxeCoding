/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReader(rd io.Reader) *Reader
 ------------------------------------------------------------------------------------
 **Description:
 NewReader returns a new Reader whose buffer has the default size.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewReader返回一个新的Reader，这个Reader的大小是默认的大小。
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	fmt.Printf("rb type: %T, len:%d\n", rb, rb.Len())
	r := bufio.NewReader(rb)
	fmt.Printf("r type: %T\n", r)
	var buf [128]byte
	n, _ := r.Read(buf[:])
	fmt.Println(string(buf[:n]))
}
