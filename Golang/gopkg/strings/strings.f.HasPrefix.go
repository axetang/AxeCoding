/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.HasPrefix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func HasPrefix(s, prefix string) bool
------------------------------------------------------------------------------------
**Description:
HasPrefix tests whether the string s begins with prefix.
------------------------------------------------------------------------------------
**要点总结:
1. HasPrefix函数检测参数字符串s是否包含前缀字符串prefix；
2. 如果prefix字符串为空，则返回true。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}
