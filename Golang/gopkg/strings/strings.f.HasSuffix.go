/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.HasSuffix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func HasSuffix(s, suffix string) bool
------------------------------------------------------------------------------------
**Description:
HasSuffix tests whether the string s ends with suffix.
------------------------------------------------------------------------------------
**要点总结:
1. HasSuffix函数检测参数字符串s是否包含前缀字符串suffix；
2. 如果suffix字符串为空，则返回true。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}
