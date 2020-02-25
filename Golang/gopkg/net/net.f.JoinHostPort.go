/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.JoinHostPort
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func JoinHostPort(host, port string) string
 ------------------------------------------------------------------------------------
 **Description:
 JoinHostPort combines host and port into a network address of the form "host:port".
 If host contains a colon, as found in literal IPv6 addresses, then JoinHostPort
 returns "[host]:port".
 See func Dial for a description of the host and port parameters.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 函数将host和port合并为一个网络地址。一般格式为"host:port"；如果host含有冒号或百分号，格式为
 "[host]:port"；
 2. 返回主机名与端口号结合后的字符串,类似host:port，比如:
	host 为 gocn_server
	port 为 8080
	则返回字符串是 gocn_server:8080
 如果主机名中存在冒号，比如:
	host 为 gocn:server
	port 为 8080
	则返回的字符串是 [gocn:server]:8080
 3. 参看Dial函数了解host和port参数描述。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr1 := net.JoinHostPort("gocn_server", "8080")
	addr2 := net.JoinHostPort("gocn:server", "8080")
	fmt.Println(addr1) // 打印结果应该为 gocn_server:8080
	fmt.Println(addr2) // 打印结果应该为 [gocn:server]:8080
}
