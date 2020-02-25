/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.PathEscape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func PathEscape(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 PathEscape escapes the string so it can be safely placed inside a URL path segment.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. PathEscape转义参数字符串s string，使得它能被安全的使用在URL的path段。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	pUrl := url.PathEscape("http://www.golang.org/1A&2")
	fmt.Println(pUrl)
	pUrl = url.PathEscape("/1A&2")
	fmt.Println(pUrl)
	pUrl = url.PathEscape("1A&2?a=1/&b=2")
	fmt.Println(pUrl)
}
