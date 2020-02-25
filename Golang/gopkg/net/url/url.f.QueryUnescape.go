/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.QueryUnescape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func QueryUnescape(s string) (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 QueryUnescape does the inverse transformation of QueryEscape, converting each 3-byte
 encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an
 error if any % is not followed by two hexadecimal digits.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. QueryUnescape函数用于将QueryEscape转码的字符串还原。它会把%AB改为字节0xAB，将'+'改为' '；
 2. 如果有某个%后面未跟两个十六进制数字，本函数会返回错误。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	uUrl, err := url.QueryUnescape("http%3A%2F%2Fwww.golang.org")
	if err == nil {
		fmt.Println(uUrl)
	} else {
		fmt.Println(err)
	}
}
