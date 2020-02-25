/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Clean
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Clean(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 Clean returns the shortest path name equivalent to path by purely lexical
 processing. It applies the following rules iteratively until no further
 processing can be done:
 1. Replace multiple Separator elements with a single one.
 2. Eliminate each . path name element (the current directory).
 3. Eliminate each inner .. path name element (the parent directory)
    along with the non-.. element that precedes it.
 4. Eliminate .. elements that begin a rooted path:
    that is, replace "/.." by "/" at the beginning of a path,
    assuming Separator is '/'.
 The returned path ends in a slash only if it represents a root directory, such
 as "/" on Unix or `C:\` on Windows.
 Finally, any occurrences of slash are replaced by Separator.
 If the result of this process is an empty string, Clean returns the string ".".
 See also Rob Pike, “Lexical File Names in Plan 9 or Getting Dot-Dot Right,”
 https://9p.io/sys/doc/lexnames.html
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Clean函数和path.Clean函数用法相同，参见package path。
*************************************************************************************/
package main

func main() {
}
