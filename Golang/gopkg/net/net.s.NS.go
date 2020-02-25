/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.NS
 **Type: struct
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
 参见LookupNS的go文件
*************************************************************************************/
package main

func main() {
}
