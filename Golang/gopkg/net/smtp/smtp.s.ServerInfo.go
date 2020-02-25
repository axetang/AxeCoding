/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: smtp
 **Element: smtp.ServerInfo
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ServerInfo struct {
     Name string   // SMTP server name
     TLS  bool     // using TLS, with valid certificate for Name
     Auth []string // advertised authentication mechanisms
 }
 ------------------------------------------------------------------------------------
 **Description:
 ServerInfo records information about an SMTP server.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ServerInfo结构体记录一个SMTP服务器的信息。
*************************************************************************************/
package main

func main() {
}
