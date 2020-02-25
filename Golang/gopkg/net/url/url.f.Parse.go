/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.Parse
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Parse(rawurl string) (*URL, error)
 ------------------------------------------------------------------------------------
 **Description:
 Parse parses rawurl into a URL structure.
 The rawurl may be relative (a path, without a host) or absolute (starting with a
 scheme). Trying to parse a hostname and path without a scheme is invalid but may
 not necessarily return an error, due to parsing ambiguities.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Parse函数解析rawurl为一个URL结构体，rawurl可以是绝对地址，也可以是相对地址；
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}
