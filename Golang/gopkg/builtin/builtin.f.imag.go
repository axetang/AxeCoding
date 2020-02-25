/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.imag
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func imag(c ComplexType) FloatType
 ------------------------------------------------------------------------------------
 **Description:
 The imag built-in function returns the imaginary part of the complex number c. The
 return value will be floating point type corresponding to the type of c.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. imag 内建函数返回复数 c 的虚部。 返回值类型与 c 的浮点类型对应。
*************************************************************************************/
package main

import "fmt"

func main() {
	fmt.Println(imag(10 + 2i))
	fmt.Println(imag(10.9 + 2.11i))
}
