/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Replace
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Replace(s, old, new string, n int) string
------------------------------------------------------------------------------------
**Description:
Replace returns a copy of the string s with the first n non-overlapping instances
of old replaced by new. If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
If n < 0, there is no limit on the number of replacements.
------------------------------------------------------------------------------------
**要点总结:
1. Replace方法返回参数s字符串的一个拷贝，s字符串的前n个出现的old字符串被new字符串替换；
2. 如果old字符串为空，替换方式需要再仔细看下总结出来，如下程序；
3. 如果n<0，则不限制替换次数，也就是全部替换。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "", "ky", -1))

	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
