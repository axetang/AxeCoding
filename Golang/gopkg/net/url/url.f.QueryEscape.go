/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.QueryEscape
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func QueryEscape(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 QueryEscape escapes the string so it can be safely placed inside a URL query.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. QueryEscape函数对s进行转码使之可以安全的用在URL查询;
 2. 所谓转码就是处理URL中的特殊字符，如“//,:,?,%”等;
 3. 转码后的%符号后要跟两个16进制字符，这是XX标准要求。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	sUrl := url.QueryEscape("http://www.golang.org/1A&2?a=1")
	fmt.Println(sUrl)
}
