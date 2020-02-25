/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Expand
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Expand(s string, mapping func(string) string) string
 ------------------------------------------------------------------------------------
 **Description:
 Expand replaces ${var} or $var in the string based on the mapping function. For
 example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 这个函数主要是把一个字符串里带$var或${var}这样的字符串按自定义的规则即函数参数mapping确定的规则
 替换成自己想要的字符串；
 2. 如果特殊字符串没有对应的替换，则替换成空。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	mapping := func(s string) string {
		m := map[string]string{"hello": "world", "go": "perfect program language"}
		return m[s]
	}
	str := "Golang is$not a $go in the ${hello}!"
	fmt.Printf("%s\n", os.Expand(str, mapping))
}
