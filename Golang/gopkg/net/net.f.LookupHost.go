/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupHost
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupHost(host string) (addrs []string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupHost looks up the given host using the local resolver. It returns a slice of
 that host's addresses.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupHost函数查询主机的网络地址序列。
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

	fmt.Println("\n\n----interfaces,lookupaddr----")
	test()
}

func test() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println("net.InterfaceAddr():", addr) //[127.0.0.1/8 10.236.15.24/24 ::1/128 fe80::3617:ebff:febe:f123/64],本地地址,ipv4和ipv6地址,这些信息可通过ifconfig命令看到
	interfaces, _ := net.Interfaces()
	fmt.Println("net.Interfaces():", interfaces) //[{1 65536 lo  up|loopback} {2 1500 eth0 34:17:eb:be:f1:23 up|broadcast|multicast}] 类型:MTU(最大传输单元),网络接口名,支持状态
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println("net.JoinHostPort():", hp) //127.0.0.1:8080,根据ip和端口组成一个addr字符串表示
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println("net.LookupAddr(\"127.0.0.1\"):", lt) //[localhost],根据地址查找到改地址的一个映射列表
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println("LookupCName(\"www.baidu.com\")", cname) //www.a.shifen.com,查找规范的dns主机名字
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println("net.LookupHost(\"www.baiduc.com\")", host) //[111.13.100.92 111.13.100.91],查找给定域名的host名称
	ip, _ := net.LookupIP("www.baidu.com")
	fmt.Println("net.LooupIP(\"www.baidu.com\")", ip) //[111.13.100.92 111.13.100.91],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.
}
