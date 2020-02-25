/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: html
 **Element: html.EscapeString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func EscapeString(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 EscapeString escapes special characters like "<" to become "&lt;". It escapes only
 five such characters: <, >, &, ' and ". UnescapeString(EscapeString(s)) == s always
 holds, but the converse isn't always true.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. EscapeString用于将特殊字符转移为html实体,如把`<`转义成`&lt;`
 2. 它只会转义下列五种字符: `<` `>` `&` `'`  `"`
 3. 需要注意的是 `UnescapeString(EscapeString(s)) == s` 返回结果一定是true,但是反过来就
 不一定是true了。
*************************************************************************************/
package main

import (
	"fmt"
	"html"
)

func main() {
	var s string = "<script>alert('xss')</script>"
	fmt.Println(html.EscapeString(s))
}
