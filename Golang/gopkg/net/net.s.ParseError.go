/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ParseError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ParseError struct {
     // Type is the type of string that was expected, such as
     // "IP address", "CIDR address".
     Type string
     // Text is the malformed text string.
     Text string
 }
 func (e *ParseError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 A ParseError is the error type of literal network address parsers.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. ParseError结构体代表一个格式错误的字符串，Type为期望的格式；
 1. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	e := net.ParseError{Type: "IP address", Text: "address malformed"}
	fmt.Println(e.Error())
}

//Output:
//invalid IP address: address malformed
