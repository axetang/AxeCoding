/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.LookupTxt
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupTXT(name string) ([]string, error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupTXT returns the DNS TXT records for the given domain name.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookupTXT函数返回指定主机的DNS TXT记录。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	txt, _ := net.LookupTXT("huawei.com")
	fmt.Println(txt)

	txts, err := net.LookupTXT("163.com")
	if err != nil {
		fmt.Println(err)
	}
	for _, txt := range txts {
		fmt.Println(txt) // 返回类似 v=spf1 include:spf.163.com -all
	}
}
