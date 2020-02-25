/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.ExpandEnv
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ExpandEnv(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 ExpandEnv replaces ${var} or $var in the string according to the values of the
 current environment variables. References to undefined variables are replaced by
 the empty string.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 这个函数主要是把一个字符串里带$var或${var}这样的字符串替换成当前环境变量中的内容；
 2. 如果没有对应的，则替换成空。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	s := "The language path is: $LANG"
	fmt.Printf("%s\n", os.ExpandEnv(s))
}
