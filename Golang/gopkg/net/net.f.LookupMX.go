/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupMX
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 type MX struct {
     Host string
     Pref uint16
 }
 func LookupMX(name string) ([]*MX, error)
 ------------------------------------------------------------------------------------
 **Description:
 An MX represents a single DNS MX record.
 LookupMX returns the DNS MX records for the given domain name sorted by preference.
 ------------------------------------------------------------------------------------
 **要点总结:
 LookupMX函数返回指定主机的按Pref字段排好序的DNS MX记录。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	mxs, err := net.LookupMX("163.com")
	if err != nil {
		fmt.Println(err)
	}
	for _, mx := range mxs {
		fmt.Println(mx)
		// 打印结果应该类似
		// &{163mx02.mxmail.netease.com. 10}
		// &{163mx01.mxmail.netease.com. 10}
		// &{163mx03.mxmail.netease.com. 10}
		// &{163mx00.mxmail.netease.com. 50}
	}
}
