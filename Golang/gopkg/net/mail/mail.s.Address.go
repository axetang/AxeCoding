/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.Address
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Address struct {
     Name    string // Proper name; may be empty.
     Address string // user@domain
 }
 func (a *Address) String() string
 ------------------------------------------------------------------------------------
 **Description:
 Address represents a single mail address. An address such as "Barry Gibbs
 <bg@example.com>" is represented as Address{Name: "Barry Gibbs", Address:
 "bg@example.com"}.
 String formats the address as a valid RFC 5322 address. If the address's name
 contains non-ASCII characters the name will be rendered according to RFC 2047.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Address类型表示一个邮箱地址。例如地址"Barry Gibbs <bg@example.com>"表示为
        Address{Name: "Barry Gibbs", Address: "bg@example.com"}；
 2. String方法依照rfc5322格式化邮件地址并返回字符串；
 3. 如果地址的Name包含非ascii字符，name会参照rfc2047格式化展示。
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
		fmt.Println(v.Name, v.Address, "String()", v.String())
	}

}
