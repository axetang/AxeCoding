/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.SimpleFold
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SimpleFold(r rune) rune
 ------------------------------------------------------------------------------------
 **Description:
 SimpleFold iterates over Unicode code points equivalent under the Unicode-defined
 simple case folding. Among the code points equivalent to rune (including rune
 itself), SimpleFold returns the smallest rune > r if one exists, or else the
 smallest rune >= 0. If r is not a valid Unicode code point, SimpleFold(r) returns r.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SimpleFold函数遍历Unicode，查找“下一个”与r相当的字符；
 2. “下一个”的意思是：码点值比r大且最靠近r的字符。如果没有“下一个”字符，则从头开始找与r相当的字符；
 3. “相当”的意思是：同一个字符在不同情况下有不同的写法，这些不同写法的字符是相当的；
 4. 这个函数通过查询caseOrbit表实现。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('T'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'
}
