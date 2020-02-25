/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Pt
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Pt(X, Y int) Point
 ------------------------------------------------------------------------------------
 **Description:
 Pt is shorthand for Point{X, Y}.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Pt函数返回Point{X , Y}
*************************************************************************************/
package main

import (
	"fmt"
	"image"
)

func main() {
	pt := image.Pt(0, 0)
	fmt.Println(pt)
}