/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.ParseError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ParseError struct {
     Layout     string
     Value      string
     LayoutElem string
     ValueElem  string
     Message    string
 }
 func (e *ParseError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 ParseError describes a problem parsing a time string.
 Error returns the string representation of a ParseError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseError结构体秒数在parsing一个时间字符串时出现的问题；
 2. Error方法返回ParseError的错误字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	pe := &time.ParseError{
		Layout:     "layout",
		Value:      "value",
		LayoutElem: "layoutelem1,layoutelem2",
		ValueElem:  "valueelem1,valueelem2",
		Message:    "Parse Error Message",
	}
	fmt.Println(pe.Error())
}
