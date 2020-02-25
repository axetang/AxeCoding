/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.ToLower,strings.ToLowerSpecial,strings.ToTitle,
strings.ToTitleSpecial,strings.ToUpper,strings.ToUpperSpecial
**Type: func
------------------------------------------------------------------------------------
**Definition:
func ToLower(s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
------------------------------------------------------------------------------------
**Description:
ToLower returns a copy of the string s with all Unicode letters mapped to their
lower case.
ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to
their lower case, giving priority to the special casing rules.
ToTitle returns a copy of the string s with all Unicode letters mapped to their title
case.
ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to
their title case, giving priority to the special casing rules.
ToUpper returns a copy of the string s with all Unicode letters mapped to their upper
case.
ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to
their upper case, giving priority to the special casing rules.
------------------------------------------------------------------------------------
**要点总结:
1. ToLower函数将字符串s修改为全部小写后返回；
2. ToLowerSpecial函数将字符串s按照参数c指定的Unicode SpecialCase小写类型后返回；
3. ToTitle将s所有字符修改为Unicode的Title case后返回；
4. ToTitleSpecial将所有字符按照c指定的Unicode SpecialCase的Title类型修改后返回；
5. ToUpper函数将字符串s全部修改为大写后返回；
6. ToUpperSpecial将字符串s按照c指定的Unicode SpecialCase的大写类型修改后返回。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	fmt.Println(strings.ToUpper("Gopher"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}
