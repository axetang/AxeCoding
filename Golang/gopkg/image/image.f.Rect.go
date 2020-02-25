/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Rect
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Rect(x0, y0, x1, y1 int) Rectangle
 ------------------------------------------------------------------------------------
 **Description:
 Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned rectangle
 has minimum and maximum coordinates swapped if necessary so that it is well-formed.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"image"
)

func main() {
	r := image.Rect(10, 10, 50, 50)
	fmt.Println(r)
}
