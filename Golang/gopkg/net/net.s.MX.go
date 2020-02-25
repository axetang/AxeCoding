/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.MX
 **Type: struct
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
 参见函数LookupMX的go文件
*************************************************************************************/
package main

func main() {
}
