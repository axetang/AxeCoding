/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: html
 **Element: html.UnescapeString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func UnescapeString(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a larger
 range of entities than EscapeString escapes. For example, "&aacute;" unescapes to
 "á", as does "&#225;" and "&#xE1;". UnescapeString(EscapeString(s)) == s always
 holds, but the converse isn't always true.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnescapeString用于将html实体转义回特殊字符,如把`&lt;`转义成`<`；
 2. 此函数可以反转义的html实体比EscapeString可转义的更多,EscapeString只会转义下列五种字符:
 `<` `>` `&` `'`  `"`。例如,`&aacute;` `&#225;` `&nbsp;` 这些都可以被反转义为特殊字符
 3. 需要注意的是 `UnescapeString(EscapeString(s)) == s` 返回结果一定是true,
 但是反过来就不一定是true了。
*************************************************************************************/
package main

import (
	"fmt"
	"html"
)

func main() {
	var s string = "&#225;&nbsp;&#225;&nbsp;&#225;"
	fmt.Println(html.UnescapeString(s))
}
