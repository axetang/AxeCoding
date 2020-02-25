/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.real
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func real(c ComplexType) FloatType
 ------------------------------------------------------------------------------------
 **Description:
 The real built-in function returns the real part of the complex number c. The return
 value will be floating point type corresponding to the type of c.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import "fmt"

func main() {
	fmt.Println(real(10 + 2i))
}
