/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.Is,unicode.IsControl,unicode.IsDigit,unicode.IsGraphic
 unicode.IsLetter,unicode.IsLower,unicode.IsMark,unicode.IsNumber,unicode.IsOneOf,
 unicode.IsPrint,unicode.IsPunct,unicode.IsSpace,unicode.IsSymbol,unicode.IsTitle,
 unicode.IsUpper
 ------------------------------------------------------------------------------------
 **Definition:
 func Is(rangeTab *RangeTable, r rune) bool
 func IsControl(r rune) bool
 func IsDigit(r rune) bool
 func IsGraphic(r rune) bool
 func IsLetter(r rune) bool
 func IsLower(r rune) bool
 func IsMark(r rune) bool
 func IsNumber(r rune) bool
 func IsOneOf(ranges []*RangeTable, r rune) bool
 func IsPrint(r rune) bool
 func IsPunct(r rune) bool
 func IsSpace(r rune) bool
 func IsSymbol(r rune) bool
 func IsTitle(r rune) bool
 func IsUpper(r rune) bool
 ------------------------------------------------------------------------------------
 **Description:
 Is reports whether the rune is in the specified table of ranges.
 IsControl reports whether the rune is a control character. The C (Other) Unicode
 category includes more code points such as surrogates; use Is(C, r) to test for
 them.
 IsDigit reports whether the rune is a decimal digit.
 IsGraphic reports whether the rune is defined as a Graphic by Unicode. Such
 characters include letters, marks, numbers, punctuation, symbols, and spaces,
 from categories L, M, N, P, S, Zs.
 IsLetter reports whether the rune is a letter (category L).
 IsLower reports whether the rune is a lower case letter.
 IsMark reports whether the rune is a mark character (category M).
 IsNumber reports whether the rune is a number (category N).
 IsOneOf reports whether the rune is a member of one of the ranges. The function
 "In" provides a nicer signature and should be used in preference to IsOneOf.
 IsPrint reports whether the rune is defined as printable by Go. Such characters
 include letters, marks, numbers, punctuation, symbols, and the ASCII space
 character, from categories L, M, N, P, S and the ASCII space character. This
 categorization is the same as IsGraphic except that the only spacing character
 is ASCII space, U+0020.
 IsPunct reports whether the rune is a Unicode punctuation character (category P).
 IsSpace reports whether the rune is a space character as defined by Unicode's White
 Space property; in the Latin-1 space this is
 '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
 Other definitions of spacing characters are set by category Z and property
 Pattern_White_Space.
 IsSymbol reports whether the rune is a symbolic character.
 IsTitle reports whether the rune is a title case letter.
 IsUpper reports whether the rune is an upper case letter.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. unicode包含了对rune的判断。这个包把所有unicode涉及到的编码进行了分类，使用结构
 type RangeTable struct {
    R16         []Range16
    R32         []Range32
    LatinOffset int
 } 来表示这个功能的字符集。这些字符集都集中列表在table.go这个源码里面;
 1. func IsControl(r rune) bool  //是否控制字符
 2. func IsDigit(r rune) bool  //是否阿拉伯数字字符，即1-9
 3. func IsGraphic(r rune) bool //是否图形字符
 4. func IsLetter(r rune) bool //是否字母
 5. func IsLower(r rune) bool //是否小写字符
 6. func IsMark(r rune) bool //是否符号字符
 7. func IsNumber(r rune) bool //是否数字字符，比如罗马数字Ⅷ也是数字字符
 8. func IsOneOf(ranges []*RangeTable, r rune) bool //是否是RangeTable中的一个
 9. func IsPrint(r rune) bool //是否可打印字符
 10. func IsPunct(r rune) bool //是否标点符号
 11. func IsSpace(r rune) bool //是否空格
 12. func IsSymbol(r rune) bool //是否符号字符
 13. func IsTitle(r rune) bool //是否title case
 14. func IsUpper(r rune) bool //是否大写字符
 15. Is函数判断字符r rune是否属于参数rangeTab *RangeTable的编码类型，可用的RangeTable参见
 go/src/unicode/tables.go。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

func TestIs() {
	s := "Hello 世界！"
	for _, r := range s {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c\n", r) // 世界
		}
	}

}

func main() {
	fmt.Println("-----Is-----")
	TestIs()
	fmt.Println("-----------------")
	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q (%d bytes):\n", c, len(string(c)))
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}
}
