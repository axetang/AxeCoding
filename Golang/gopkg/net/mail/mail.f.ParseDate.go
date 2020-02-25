/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.ParseDate
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseDate(date string) (time.Time, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseDate parses an RFC 5322 date string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 解析date字符串，返回time.Time。
*************************************************************************************/
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	t, _ := mail.ParseDate("Mon, 23 Jun 2015 11:40:36 -0400")
	fmt.Println(t)
}
