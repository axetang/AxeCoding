/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func complex(r, i FloatType) ComplexType
 ------------------------------------------------------------------------------------
 **Description:
 The complex built-in function constructs a complex value from two floating-point
 values. The real and imaginary parts must be of the same size, either float32 or
 float64 (or assignable to them), and the return value will be the corresponding
 complex type (complex64 for float32, complex128 for float64).
 ------------------------------------------------------------------------------------
 **要点总结:
 complex内建函数将两个浮点数值构造成一个复数值。其实部和虚部的大小必须相同，即float32或float64，
 其返回值即为对应的复数类型（complex64对应float32，complex128对应float64）。
*************************************************************************************/
package main

import "fmt"

func main() {
	fmt.Println(complex(10, 2))
}
