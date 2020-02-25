/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package:
 **Element:
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type ByteWriter interface {
    WriteByte(c byte) error
 }
 ------------------------------------------------------------------------------------
 **Description:
 ByteWriter is the interface that wraps the WriteByte method.
 ------------------------------------------------------------------------------------
 **要点总结:1. 待补充
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
