/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.print
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func print(args ...Type)
 ------------------------------------------------------------------------------------
 **Description:
 The print built-in function formats its arguments in an implementation-specific way
 and writes the result to standard error. Print is useful for bootstrapping and
 debugging; it is not guaranteed to stay in the language.
 ------------------------------------------------------------------------------------
 **要点总结:
 向标准输出打印。
*************************************************************************************/
package main

func main() {
	print(1, 2, 3, "\n")
	println(4.5, 6, complex(1, 2), "hhh")
}
