/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.Header
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Message struct {
     Header Header
     Body   io.Reader
 }
 ------------------------------------------------------------------------------------
 **Description:
 A Message represents a parsed mail message.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

func main() {
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>, Axe <axetang@163.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

axe
Message body
ttt
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	fmt.Println("------AddressList,Date,-------------")
	f, _ := header.AddressList("From")
	fmt.Println("AddressList(From)", f)
	d, _ := header.Date()
	fmt.Println("Date:", d)

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("----Body----------\n%s\n", body)
}
