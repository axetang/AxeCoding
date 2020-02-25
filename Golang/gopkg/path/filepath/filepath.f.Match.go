/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Match
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Match(pattern, name string) (matched bool, err error)
 -----------------------------------------------------------------------------------
 **Description:
 Match reports whether name matches the shell file name pattern. The pattern
 syntax is:
 pattern:
 	{ term }
 term:
 	'*'         matches any sequence of non-Separator characters
 	'?'         matches any single non-Separator character
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
 On Windows, escaping is disabled. Instead, '\\' is treated as path separator.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Match函数的使用和功能与path.Match一致，参看package path。
 2. Pattern规则如上。
*************************************************************************************/
package main

func main() {
}
