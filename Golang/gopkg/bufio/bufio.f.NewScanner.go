/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewScanner
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewScanner(r io.Reader) *Scanner
 ------------------------------------------------------------------------------------
 **Description:
 NewScanner returns a new Scanner to read from r. The split function defaults to
 ScanLines.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewScanner函数定义如下：
 func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:            r,
		split:        ScanLines,
		maxTokenSize: MaxScanTokenSize,
	}
 }
 2. NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。
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
	s := bufio.NewScanner(r)
	fmt.Printf("s type: %T\n", s)
}
