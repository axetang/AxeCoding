/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupAddr
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupAddr(addr string) (names []string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupAddr performs a reverse lookup for the given address, returning a list of names
 mapping to that address.
 When using the host C library resolver, at most one result will be returned. To
 bypass the host resolver, use a custom Resolver.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupAddr查询某个地址，返回映射到该地址的主机名序列，本函数和LookupHost不互为反函数。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.JoinHostPort("163.com", "8080")
	fmt.Println(addr) // 打印结果应该为 gocn_server:8080
}
