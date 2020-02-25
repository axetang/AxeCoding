/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.AddrError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type AddrError struct {
     Err  string
     Addr string
 }
 func (e *AddrError) Error() string
 func (e *AddrError) Temporary() bool
 func (e *AddrError) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 0. AddrError代表地址信息的错误；
 1. Error方法返回错误信息字符串；
 2. Temporary和Timeout方法判断错误性质.
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	err := net.AddrError{Err: "this is DNS Config Error", Addr: "1.1.1.1"}
	fmt.Println(err.Error())
	fmt.Println(err.Temporary())
	fmt.Println(err.Timeout())

}
