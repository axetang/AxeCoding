/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.ParseAddressList
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseAddressList(list string) ([]*Address, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseAddressList parses the given string as a list of addresses.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseAddressList函数将参数list string作为一串邮箱地址解析；
 2. 返回[]*Address。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {

	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}
}
