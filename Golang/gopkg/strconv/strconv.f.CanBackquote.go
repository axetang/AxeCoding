/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.CanBackquote
 **Type: func
 **Note: this code segment was edited in Hongkong Univ. on Feb. 15, 2019.
 ------------------------------------------------------------------------------------
 **Definition:
 func CanBackquote(s string) bool
 ------------------------------------------------------------------------------------
 **Description:a
 CanBackquote reports whether the string s can be represented unchanged as a
 single-line backquoted string without control characters other than tab.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Backquote：反引号
 2. 如果字符串参数s包含反引号或者"\t"以外的控制字符，则CanBackquote返回false；
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
	fmt.Println(strconv.CanBackquote("can backquote\t  this"))
	fmt.Println(strconv.CanBackquote("can backquote…… ^  this"))
	fmt.Println(strconv.CanBackquote("can backquote\t \nthis"))
}
