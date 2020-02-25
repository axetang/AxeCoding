/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.ByteReader
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type ByteReader interface {
    ReadByte() (byte, error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 ByteReader is the interface that wraps the ReadByte method. ReadByte reads and
 returns the next byte from the input or any error encountered. If ReadByte
 returns an error, no input byte was consumed, and the returned byte value is
 undefined.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ByteReader接口封装了ReadByte方法，返回读出的一个字节和可能的错误；
 2. 需要结合其他函数一起来看这个接口的使用，后续再补充内容
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
