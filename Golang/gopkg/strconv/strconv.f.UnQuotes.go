/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.Unquote,strconv.UnquoteChar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Unquote(s string) (string, error)
 func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Unquote interprets s as a single-quoted, double-quoted, or backquoted Go string
 literal, returning the string value that s quotes. (If s is single-quoted, it would
 be a Go character literal; Unquote returns the corresponding one-character string.)
 UnquoteChar decodes the first character or byte in the escaped string or character
 literal represented by the string s. It returns four values:
	 1) value, the decoded Unicode code point or byte value;
	 2) multibyte, a boolean indicating whether the decoded character requires a
	 multibyte UTF-8 representation;
	 3) tail, the remainder of the string after the character; and
	 4) an error that will be nil if the character is syntactically valid.
 The second argument, quote, specifies the type of literal being parsed and therefore
 which escaped quote character is permitted. If set to a single quote, it permits
 the sequence \' and disallows unescaped '. If set to a double quote, it permits \"
 and disallows unescaped ". If set to zero, it does not permit either escape and
 allows both quote characters to appear unescaped.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Unquote去掉字符串首尾的双引号对、单引号对或反引号对；如果不是成对出现的引号，则会返回错误；
 2. UnquoteChar返回四个值：
	1）value rune: 被解码的unicode码点；
	2）mutibyte bool：被解码的unicode码点字符是否需要多字符；
	3）字符串s去除被decode的字符的新字符串；
	4）err error：后续更新。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s, err := strconv.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	//s, err = strconv.Unquote("\"The string must be either double-quoted")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err)

	fmt.Println("---UnquoteChar---")
	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)
}
