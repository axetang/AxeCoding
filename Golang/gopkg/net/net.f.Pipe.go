/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Pipe
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Pipe() (Conn, Conn)
 ------------------------------------------------------------------------------------
 **Description:
 Pipe creates a synchronous, in-memory, full duplex network connection; both ends
 implement the Conn interface. Reads on one end are matched with writes on the other,
 copying data directly between the two; there is no internal buffering.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Pipe创建一个内存中的同步、全双工网络连接。连接的两端都实现了Conn接口。一端的读取对应另一端的
 写入，直接将数据在两端之间作拷贝；没有内部缓冲。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	c1, c2 := net.Pipe()
	fmt.Println(c1)
	fmt.Println(c2)
}
