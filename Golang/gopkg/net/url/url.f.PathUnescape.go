/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.PathUnescape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func PathUnescape(s string) (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 PathUnescape does the inverse transformation of PathEscape, converting each 3-byte
 encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an
 error if any % is not followed by two hexadecimal digits.
 PathUnescape is identical to QueryUnescape except that it does not unescape '+'
 to ' ' (space).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. PathUnescape和PathEscape函数对应，反转义字符串s为其原始形式；
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	pUrl, err := url.PathUnescape("http%3A%2F%2Fwww.golang.org")
	if err == nil {
		fmt.Println(pUrl)
	} else {
		fmt.Println(err)
	}
	pUrl, _ = url.PathUnescape("%2F1A&2")
	fmt.Println(pUrl)
	pUrl, _ = url.PathUnescape("1A&2%3Fa=1%2F&b=2")
	fmt.Println(pUrl)

}
