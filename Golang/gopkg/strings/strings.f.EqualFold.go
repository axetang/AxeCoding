/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.EqualFold
**Type: func
------------------------------------------------------------------------------------
**Definition:
func EqualFold(s, t string) bool
------------------------------------------------------------------------------------
**Description:
EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under
Unicode case-folding.
------------------------------------------------------------------------------------
**要点总结:
EqualFold函数判断两个UTF-8编码字符串参数s和t在不区分大小写情况下是否相同。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Go", "gO"))
	fmt.Println(strings.EqualFold("Axe Tang", "aXe tANg"))
}
