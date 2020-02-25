/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.SpecialCase
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type SpecialCase []CaseRange
 var AzeriCase SpecialCase = _TurkishCase
 var TurkishCase SpecialCase = _TurkishCase
 func (special SpecialCase) ToLower(r rune) rune
 func (special SpecialCase) ToTitle(r rune) rune
 func (special SpecialCase) ToUpper(r rune) rune
 ------------------------------------------------------------------------------------
 **Description:
 SpecialCase represents language-specific case mappings such as Turkish. Methods of
 SpecialCase customize (by overriding) the standard mappings.
 ToLower maps the rune to lower case giving priority to the special mapping.
 ToTitle maps the rune to title case giving priority to the special mapping.
 ToUpper maps the rune to upper case giving priority to the special mapping.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. SpecialCase代表指定语言的case mapping；
 1. ToLower将字符r rune按照指定的special mapping转换后返回；
 2. ToUpper将字符r rune按照指定的special mapping转换后返回；
 3. ToTitle将字符r rune按照指定的special mapping转换后返回。
 4. 后续更新理解。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//t := unicode.TurkishCase
	t := unicode.AzeriCase
	const lci = 'i'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToTitle(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))
	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))
	println()

}
