/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Repeat
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Repeat(s string, count int) string
------------------------------------------------------------------------------------
**Description:
Repeat returns a new string consisting of count copies of the string s.
It panics if count is negative or if the result of (len(s) * count) overflows.
------------------------------------------------------------------------------------
**要点总结:
1. Repeat方法返回由参数s字符串重复count次后的新字符串；
2. 如果count是负值，或者len(s)*count溢出，则函数抛出panic。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
}
