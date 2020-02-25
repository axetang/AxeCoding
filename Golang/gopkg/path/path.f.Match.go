/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Match
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Match(pattern, name string) (matched bool, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Match reports whether name matches the shell pattern. The pattern syntax is:
 pattern:
 	{ term }
 term:
 	'*'         matches any sequence of non-/ characters
 	'?'         matches any single non-/ character
 	'[' [ '^' ] { character-range } ']'
 	            character class (must be non-empty)
 	c           matches character c (c != '*', '?', '\\', '[')
 	'\\' c      matches character c
 character-range:
 	c           matches character c (c != '\\', '-', ']')
 	'\\' c      matches character c
 	lo '-' hi   matches character c for lo <= c <= hi
 Match requires pattern to match all of name, not just a substring. The only
 possible returned error is ErrBadPattern, when pattern is malformed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Match函数用于文件名匹配，只有完全匹配规则要求才返回true，nil
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("*", "alll"))  //true nil
	fmt.Println(path.Match("*", "a/lll")) //false nil
	fmt.Println(path.Match("?", "alll"))  //false nil
	fmt.Println(path.Match("?", "a"))     //true nil
	fmt.Println(path.Match("a", "a"))     //true nil
	fmt.Println(path.Match("\\a", "a"))   //true nil
}
