/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Resolver
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupSRV tries to resolve an SRV query of the given service, protocol, and domain
 name. The proto is "tcp" or "udp". The returned records are sorted by priority and
 randomized by weight within a priority.
 LookupSRV constructs the DNS name to look up following RFC 2782. That is, it looks
 up _service._proto.name. To accommodate services publishing SRV records under
 non-standard names, if both service and proto are empty strings, LookupSRV looks
 up name directly.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. SRV代表一条DNS SRV记录（资源记录），记录某个服务由哪台计算机提供。
 1. LookupSRV函数尝试执行指定服务、协议、主机的SRV查询;
 2. 协议proto为"tcp" 或"udp"。返回的记录按Priority字段排序，同一优先度按Weight字段随机排序;
 3. LookupSRV函数按照RFC2782的规定构建用于查询的DNS名,也就是说，它会查询_service._proto.name,
 为了适应将服务的SRV记录发布在非规范名下的情况，如果service和proto参数都是空字符串，函数会直接查询
 name。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	cname, addrs, err := net.LookupSRV("", "", "netease.com")
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, addr := range addrs {
		fmt.Println("addr:", addr) // 返回类似 &{xmpp-server.l.google.com. 5269 5 0}
	}
	fmt.Println("cname:", cname) //返回类似 _xmpp-server._tcp.gocn.im.
}
