/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Clean
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Clean(path string) string
 ------------------------------------------------------------------------------------
 **Description:
 Clean returns the shortest path name equivalent to path by purely lexical processing. It applies the following rules iteratively until no further processing can be done:
 1. Replace multiple slashes with a single slash.
 2. Eliminate each . path name element (the current directory).
 3. Eliminate each inner .. path name element (the parent directory)
    along with the non-.. element that precedes it.
 4. Eliminate .. elements that begin a rooted path:
    that is, replace "/.." by "/" at the beginning of a path.
 The returned path ends in a slash only if it is the root "/".
 If the result of this process is an empty string, Clean returns the string ".".
 See also Rob Pike, “Lexical File Names in Plan 9 or Getting Dot-Dot Right,” https://9p.io/sys/doc/lexnames.html
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Clean函数返回等价的最短路径；
 2. Clean的清理规则：
 -		1）用一个斜线替换多个斜线
 -		2）清除当前路径.
 -		3）清除内部的..和他前面的元素，如a/b/.. 得到结果a
 -		4）以/..开头的，变成/
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Clean("a/c"))               // a/c
	fmt.Println(path.Clean("a//c"))              //以一个/代替// , a/c
	fmt.Println(path.Clean("a/c/."))             //清除. , a/c
	fmt.Println(path.Clean("a/c/b/.."))          // 清除内部..以及前面的元素b; a/c
	fmt.Println(path.Clean("111/../a/c"))        // 清除内部..以及前面的元素111; a/c
	fmt.Println(path.Clean("/../a/b/../././/c")) // 清除/..开头,..以及前面的元素b; /a/c
}
