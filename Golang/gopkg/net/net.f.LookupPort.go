/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupPort
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupPort(network, service string) (port int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupPort looks up the port for the given network and service.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupPort函数查询指定网络和服务的（默认）端口;
 2. 下面程序需要更新。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	p, e := net.LookupPort("163.com", "smtp")
	if e != nil {
		fmt.Println("port:", p)

	} else {
		fmt.Println(e)
	}
}
