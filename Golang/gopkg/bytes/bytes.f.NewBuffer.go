/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bytes
 **Element: bytes.NewBuffer
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewBuffer(buf []byte) *Buffer
 ------------------------------------------------------------------------------------
 **Description:
 NewBuffer creates and initializes a new Buffer using buf as its initial contents.
 The new Buffer takes ownership of buf, and the caller should not use buf after
 this call. NewBuffer is intended to prepare a Buffer to read existing data. It
 can also be used to size the internal buffer for writing. To do that, buf should
 have the desired capacity but a length of zero.
 In most cases, new(Buffer) (or just declaring a Buffer variable) is sufficient
 to initialize a Buffer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查阅标准库源代码，NewBuffer函数的定义和实现如下
 // In most cases, new(Buffer) (or just declaring a Buffer variable) is
 // sufficient to initialize a Buffer.
 func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
 2. NewBuffer创建一个新的Buffer，并使用参数buf进行初始化。这个buf用来作为准备要读的数据；
 也可以用来指定写缓冲区的大小，这时buf应该是cap(buf)为指定大小，但是len(buf)为0。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("123456")
	fmt.Printf("buf cap: %d, len:%d\n", cap(buf), len(buf))
	b := bytes.NewBuffer(buf)
	c := *b
	fmt.Printf("b type: %T\n c type : %T\n", b, c)
	var data [6]byte
	b.Read(data[:])
	fmt.Println(string(data[:]))
}
