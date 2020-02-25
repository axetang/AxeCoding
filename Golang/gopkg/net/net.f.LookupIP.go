/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupIP
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupIP(host string) ([]IP, error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupIP looks up host using the local resolver. It returns a slice of that
 host's IPv4 and IPv6 addresses.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupIP函数查询主机的ipv4和ipv6地址序列。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	h := "www.baidu.com"
	cname, err := net.LookupCNAME(h)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cname) // 打印结果应该为 bbs.gocn.im.
	host, _ := net.LookupHost(h)
	ip, _ := net.LookupIP(h)
	fmt.Println("-----LookupHost,LookupIP--------")
	fmt.Println("host:", host)
	fmt.Println("IP:", ip)

}
