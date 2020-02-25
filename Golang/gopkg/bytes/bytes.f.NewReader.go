/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.NewReader
**Type: func
------------------------------------------------------------------------------------
**Definition:
func NewReader(b []byte) *Reader
------------------------------------------------------------------------------------
**Description:
NewReader returns a new Reader reading from b.
------------------------------------------------------------------------------------
**要点总结:
1. NewReader函数返回一个从b([]byte)读取的*Reader；
2. 查阅源码，NewReader实现如下，即使用参数b构建一个Reader并返回。
func NewReader(b []byte) *Reader { return &Reader{b, 0, -1} }
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewReader([]byte("12345"))
	fmt.Println(b.Len())
}
