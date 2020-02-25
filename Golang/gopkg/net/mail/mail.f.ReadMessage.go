/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.ReadMessage
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadMessage(r io.Reader) (msg *Message, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ReadMessage reads a message from r. The headers are parsed, and the body of the
 message will be available for reading from msg.Body.
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
From: Gopher <from@example.com>
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

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
