/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.OpError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type OpError struct {
     // Op is the operation which caused the error, such as
     // "read" or "write".
     Op  string
     // Net is the network type on which this error occurred,
     // such as "tcp" or "udp6".
     Net string
     // For operations involving a remote network connection, like
     // Dial, Read, or Write, Source is the corresponding local
     // network address.
     Source Addr
     // Addr is the network address for which this error occurred.
     // For local operations, like Listen or SetDeadline, Addr is
     // the address of the local endpoint being manipulated.
     // For operations involving a remote network connection, like
     // Dial, Read, or Write, Addr is the remote address of that
     // connection.
     Addr Addr
     // Err is the error that occurred during the operation.
     Err error
 }
 func (e *OpError) Error() string
 func (e *OpError) Temporary() bool
 func (e *OpError) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 OpError is the error type usually returned by functions in the net package. It
 describes the operation, network type, and address of an error.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. OpError是经常被net包的函数返回的错误类型。它描述了该错误的操作、网络类型和网络地址。
 1. Error方法返回错误信息字符串；
 2. Temporary和Timeout方法判断错误性质.
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	err := net.OpError{"Read", "tcp", "1.1.1.1", "2.2.2.2", "THIS IS AN OpError"}
	fmt.Println(err.Error())
	fmt.Println(err.Temporary())
	fmt.Println(err.Timeout())
	fmt.Println(err.Op, err.Net, err.Source, err.Addr, err.Err)
}
