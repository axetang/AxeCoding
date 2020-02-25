/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Join
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Join(a []string, sep string) string
------------------------------------------------------------------------------------
**Description:
Join concatenates the elements of a to create a single string. The separator string
sep is placed between elements in the resulting string.
------------------------------------------------------------------------------------
**要点总结:
1. Join方法用sep string分隔符将字符串切片a的元素链接起来形成一个新的字符串并返回。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}
