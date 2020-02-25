/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.ByteScanner
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type ByteScanner interface {
    ByteReader
    UnreadByte() error
 }
 ------------------------------------------------------------------------------------
 **Description:
 ByteScanner is the interface that adds the UnreadByte method to the basic ReadByte
 method. UnreadByte causes the next call to ReadByte to return the same byte as the
 previous call to ReadByte. It may be an error to call UnreadByte twice without an
 intervening call to ReadByte.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 待补充
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
