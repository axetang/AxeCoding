/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bytes
 **Element: bytes.NewBuffer
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewBufferString(s string) *Buffer
 ------------------------------------------------------------------------------------
 **Description:
 NewBufferString creates and initializes a new Buffer using string s as its
 initial contents. It is intended to prepare a buffer to read an existing string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查阅标准库源代码，NewBuffer函数的定义和实现如下
 // NewBufferString creates and initializes a new Buffer using string s as its
 // initial contents. It is intended to prepare a buffer to read an existing
 // string. In most cases, new(Buffer) (or just declaring a Buffer variable) is
 // sufficient to initialize a Buffer.
 func NewBufferString(s string) *Buffer {
	return &Buffer{buf: []byte(s)}
 }
 2. 函数NewBufferString将字符串参数s转换成字符数组([]bytes(s))，然后赋值给结构体Buffer的buf成员，
 并以结构体指针*Buffer返回。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("123456")
	var data [6]byte
	b.Read(data[:])
	fmt.Println(string(data[:]))
}
