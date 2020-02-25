/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.Quote,strconv.QuoteRune,strconv.QuoteRuneToASCII,
 strconv.QuoteRuneToGraphic,strconv.QuoteToASCII,strconv.QuoteToGraphic
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Quote(s string) string
 func QuoteRune(r rune) string
 func QuoteRuneToASCII(r rune) string
 func QuoteRuneToGraphic(r rune) string
 func QuoteToASCII(s string) string
 func QuoteToGraphic(s string) string
 ------------------------------------------------------------------------------------
 **Description:
 Quote returns a double-quoted Go string literal representing s. The returned string
 uses Go escape sequences (\t, \n, \xFF, \u0100) for control characters and
 non-printable characters as defined by IsPrint.
 QuoteRune returns a single-quoted Go character literal representing the rune. The
 returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for control
 characters and non-printable characters as defined by IsPrint.
 QuoteRuneToASCII returns a single-quoted Go character literal representing the rune.
 The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for non-ASCII
 characters and non-printable characters as defined by IsPrint.
 QuoteRuneToGraphic returns a single-quoted Go character literal representing the
 rune. The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
 non-ASCII characters and non-printable characters as defined by IsGraphic.
 QuoteToASCII returns a double-quoted Go string literal representing s. The returned
 string uses Go escape sequences (\t, \n, \xFF, \u0100) for non-ASCII characters and
 non-printable characters as defined by IsPrint.
 QuoteToGraphic returns a double-quoted Go string literal representing s. The
 returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for non-ASCII
 characters and non-printable characters as defined by IsGraphic.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Quote方法给参数s string包裹双引号对后返回这个新的Go字符串字面量；
 2. QuoteRune给参数r rune包裹单引号对后返回这个新的Go字符串字面量；
 3. QuoteRuneToASCII将参数r rune转换为ASCII码后包裹单引号，返回这个新的Go字符串；
 4. QuoteRuneToGraphic给参数r rune转化为Go符号并包裹单引号对后返回这个新的Go字符串；
 5. QuoteToASCII将参数s string转化为双引号包裹的asc字符串，非asc字符用go escape sequences代替；
 6. QuoteToGraphic将参数s string转化为Go符号并包裹双引号对后返回；
 7. ToGraphic类函数需要后续研究unicode包后再更新具体功能。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)
	s = strconv.Quote(`"Fran \t & Freddie's \n Diner	☺"`)
	fmt.Println(s)
	println()

	s = strconv.QuoteRune('☺')
	fmt.Println(s)
	println()

	s = strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)
	println()

	s = strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s)
	println()

	s = strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)
	println()

	s = strconv.QuoteToGraphic("☺\t")
	fmt.Println(s)
	println()
}
