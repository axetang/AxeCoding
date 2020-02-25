/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupNS
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 type NS struct {
     Host string
 }
 func LookupNS(name string) ([]*NS, error)
 ------------------------------------------------------------------------------------
 **Description:
 An NS represents a single DNS NS record.
 LookupNS returns the DNS NS records for the given domain name.
 ------------------------------------------------------------------------------------
 **要点总结:
 LookupNS函数返回指定主机的DNS NS记录。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	ns, err := net.LookupNS("163.com")
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range ns {
		fmt.Println(n)
	}
}
