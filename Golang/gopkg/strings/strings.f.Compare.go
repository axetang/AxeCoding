/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Compare
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Compare(a, b string) int
------------------------------------------------------------------------------------
**Description:
Compare returns an integer comparing two strings lexicographically. The result will
be 0 if a==b, -1 if a < b, and +1 if a > b.
Compare is included only for symmetry with package bytes. It is usually clearer and
always faster to use the built-in string comparison operators ==, <, >, and so on.
------------------------------------------------------------------------------------
**要点总结:
1. Compare函数按照词典来比较两个字符串，如果a==b返回0，如果a<b，返回-1，如果a>b，返回1；
2. Compare是为了与bytes包对称而包含进来的函数。通常使用内奸的字符串比较操作符，如==，<,>,更方便快捷。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}
