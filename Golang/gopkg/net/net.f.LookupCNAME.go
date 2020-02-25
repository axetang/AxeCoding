/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupCNAME
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupCNAME(host string) (cname string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupCNAME returns the canonical name for the given host. Callers that do not care
 about the canonical name can call LookupHost or LookupIP directly; both take care of
 resolving the canonical name as part of the lookup.
 A canonical name is the final name after following zero or more CNAME records.
 LookupCNAME does not return an error if host does not contain DNS "CNAME" records,
 as long as host resolves to address records.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupCNAME函数查询name的规范DNS名（但该域名未必可以访问）;
 2. 如果调用者不关心规范名可以直接调用LookupHost或者LookupIP；这两个函数都会在查询时考虑到规范名。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	h := "baidu.com"
	cname, err := net.LookupCNAME(h)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cname)
	host, _ := net.LookupHost(h)
	ip, _ := net.LookupIP(h)
	fmt.Println("-----LookupHost,LookupIP--------")
	fmt.Println("host:", host)
	fmt.Println("IP:", ip)

}
