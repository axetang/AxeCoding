/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.ParseAddress
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseAddress(address string) (*Address, error)
 ------------------------------------------------------------------------------------
 **Description:
 Parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseAddress函数解析一个单一的RFC5322地址；
 2. 返回*Address。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	e, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.Name, e.Address)
}
