/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Join
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Join(elem ...string) string
 -----------------------------------------------------------------------------------
 **Description:
 Join joins any number of path elements into a single path, adding a Separator
 if necessary. Join calls Clean on the result; in particular, all empty strings
 are ignored. On Windows, the result is a UNC path if and only if the first
 path element is a UNC path.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Join函数和path.Join函数使用和功能一致，参看package path.
*************************************************************************************/
package main

func main() {
}
