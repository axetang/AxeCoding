/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.println
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func println(args ...Type)
 ------------------------------------------------------------------------------------
 **Description:
 The println built-in function formats its arguments in an implementation-specific
 way and writes the result to standard error. Spaces are always added between
 arguments and a newline is appended. Println is useful for bootstrapping and
 debugging; it is not guaranteed to stay in the language.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
	print(1, 2, 3, "\n")
	println(4.5, 6, complex(1, 2), "hhh")
}
