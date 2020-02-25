/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: mail
 **Element: mail
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type AddressParser struct {
     // WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
     WordDecoder *mime.WordDecoder
 }
 func (p *AddressParser) Parse(address string) (*Address, error)
 func (p *AddressParser) ParseList(list string) ([]*Address, error)
 ------------------------------------------------------------------------------------
 **Description:
 An AddressParser is an RFC 5322 address parser.
 Parse parses a single RFC 5322 address of the form "Gogh Fir <gf@example.com>" or
 "foo@example.com".
 ParseList parses the given string as a list of comma-separated addresses of the
 form "Gogh Fir <gf@example.com>" or "foo@example.com".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. AddressParser结构体是一个基于RFC5322的地址解析器；
 2. Parse方法解析一个单一的RFC5322地址字符串，返回地址信息*Address；
 3. ParseList方法解析给定字符串list string为一个地址列表[]*Address并返回；
 4. 结构体成员WordDecoder的作用待后续更新。
*************************************************************************************/
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	ap := &mail.AddressParser{}
	ad, _ := ap.Parse("axe <axetang@163.com>")
	fmt.Println(ad)
	ad, _ = ap.Parse("axetang@163.com")
	fmt.Println(ad)

	fmt.Println("----ParseList------")
	al, _ := ap.ParseList("axe <axetang@163.com>,felix <felixtang@163.com>,tangfei@huawei.com")
	for _, a := range al {
		fmt.Println(a)
	}
}
