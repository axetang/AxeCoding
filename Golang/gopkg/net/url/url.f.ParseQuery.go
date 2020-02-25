/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.ParseQuery
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseQuery(query string) (Values, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseQuery parses the URL-encoded query string and returns a map listing the values
 specified for each key. ParseQuery always returns a non-nil map containing all the
 valid query parameters found; err describes the first decoding error encountered,
 if any.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 通过ParseQuery解析URL编码的查询字符串，并返回一个map ，列出每个键的值;
 2. ParseQuery总是返回一个非空的map，包含所有有效的查询参数，通过";"与"&"分割;如果遇到解析异常将
 通过err返回错误信息;
 3. 要了解下查询字符串结构。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	//http://127.0.0.1:6060/src/pkg/url/url.go?h=%22%26%22;s=14652:14657#L534&test=abc;123
	test, err := url.ParseQuery("h=%22%26%22;s=14652:14657#L534&test=abc;123")
	if err == nil {
		fmt.Println(test)
	} else {
		fmt.Println(err)
	}

}
