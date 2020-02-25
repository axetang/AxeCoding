/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.InvalidAddrError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type InvalidAddrError string
 func (e InvalidAddrError) Error() string
 func (e InvalidAddrError) Temporary() bool
 func (e InvalidAddrError) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 0. InvalidAddrError类型代表非法地址错误信息；
 1. Error方法返回错误信息字符串；
 2. Temporary和Timeout方法判断错误性质。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	var e net.InvalidAddrError = "this is a InvalidAddrError"
	fmt.Println(e.Error())
	fmt.Println(e.Temporary())
	fmt.Println(e.Timeout())

}
