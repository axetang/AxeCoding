/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Lookup
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Lookup(name string) *Flag
 ------------------------------------------------------------------------------------
 **Description:
 Lookup returns the Flag structure of the named command-line flag, returning
 nil if none exists.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Lookup函数返回参数name string指定的命令行参数的指针即*Flag结构体信息；
 2. 如果不存在该参数，则函数返回nil。
*************************************************************************************/
package main

func main() {
}
