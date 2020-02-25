/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Closer
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Closer interface {
    Close() error
 }
 ------------------------------------------------------------------------------------
 **Description:
 Closer is the interface that wraps the basic Close method. The behavior of Close
 after the first call is undefined. Specific implementations may document their own
 behavior.
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
