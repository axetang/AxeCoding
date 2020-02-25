/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail.Header
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Header map[string][]string
 func (h Header) AddressList(key string) ([]*Address, error)
 func (h Header) Date() (time.Time, error)
 func (h Header) Get(key string) string
 ------------------------------------------------------------------------------------
 **Description:
 A Header represents the key-value pairs in a mail message header.
 AddressList parses the named header field as a list of addresses.
 Date parses the Date header field.
 Get gets the first value associated with the given key. It is case insensitive;
 CanonicalMIMEHeaderKey is used to canonicalize the provided key. If there are no
 values associated with the key, Get returns "". To access multiple values of a key,
 or to use non-canonical keys, access the map directly.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Header代表邮件消息头的key-value对；
 1. mail.ReadMessage函数从io.Reader中读取消息内容并解析，返回*Message结构体。该结构体包含了两个
 成员：Header Header和Body io.Reader；
 2. AddressList方法把named header字段解析为邮件地址列表[]*Address；
 3. Date解析header中的Date字段，返回time.Time结构体；
 4. Get返回与给定参数key string对应的第一个值，如果找不到返回空串。
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
