/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Count
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Count(s, substr string) int
------------------------------------------------------------------------------------
**Description:
Count counts the number of non-overlapping instances of substr in s. If substr is
an empty string, Count returns 1 + the number of Unicode code points in s.
------------------------------------------------------------------------------------
**要点总结:
1. Count函数统计字符串s中出现子字符串substr次数；
2. 如果substr是空串，Count方法返回1+（s中的unicode字符数量）
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}
