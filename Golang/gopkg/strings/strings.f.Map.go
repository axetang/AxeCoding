/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Map
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Map(mapping func(rune) rune, s string) string
------------------------------------------------------------------------------------
**Description:
Map returns a copy of the string s with all its characters modified according to
the mapping function. If mapping returns a negative value, the character is dropped
from the string with no replacement.
------------------------------------------------------------------------------------
**要点总结:
1. Map方法返回参数s字符串的一个新拷贝；
2. 这个拷贝字符串中所有字符都会被依照mapping函数参数来进行修改；
3. 如果mapping函数执行返回负值，则相应字符会从字符串中去掉也不会被替换。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
