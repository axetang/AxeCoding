/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.To,unicode.ToLower,unicode.ToTitle,unicode.ToUpper
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func To(_case int, r rune) rune
 func ToLower(r rune) rune
 func ToTitle(r rune) rune
 func ToUpper(r rune) rune
 ------------------------------------------------------------------------------------
 **Description:
 To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
 ToLower maps the rune to lower case.
 ToTitle maps the rune to title case.
 ToUpper maps the rune to upper case.
  ------------------------------------------------------------------------------------
 **要点总结:
 1. ToLower将参数r rune转换为lower case后返回；
 2. ToUpper将参数r rune转换为upper case后返回；
 3. ToTitle将参数r rune转换为title case后返回；
 4. To函数依据参数_case int指定的case转换参数r rune后返回；
*************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("--------To-----------")
	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))
	fmt.Println("--------ToLower-----------")
	const ucG1 = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG1))
	fmt.Println("--------ToTitle-----------")
	const ucG2 = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG2))
	fmt.Println("--------ToUpper-----------")
	const ucG3 = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG3))
}
