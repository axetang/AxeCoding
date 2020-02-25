/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Point
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Point struct {
     X, Y int
 }
 var ZP Point
 func (p Point) Add(q Point) Point
 func (p Point) Div(k int) Point
 func (p Point) Eq(q Point) bool
 func (p Point) In(r Rectangle) bool
 func (p Point) Mod(r Rectangle) Point
 func (p Point) Mul(k int) Point
 func (p Point) String() string
 func (p Point) Sub(q Point) Point
 ------------------------------------------------------------------------------------
 **Description:
 A Point is an X, Y coordinate pair. The axes increase right and down.
 ZP is the zero Point.
 Add returns the vector p+q.
 Div returns the vector p/k.
 Eq reports whether p and q are equal.
 In reports whether p is in r.
 Mod returns the point q in r such that p.X-q.X is a multiple of r's width and
 p.Y-q.Y is a multiple of r's height.
 Mul returns the vector p*k.
 String returns a string representation of p like "(3,4)".
 Sub returns the vector p-q.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Point结构体是X, Y坐标对。坐标轴是向右（X）向下（Y）的。既可以表示点，也可以表示向量；
 1. ZP是原点；
 2. Mod方法返回r范围内的某点q，满足p.X-q.X是r宽度的倍数，p.Y-q.Y是r高度的倍数；
 3. Add返回p+q；
 4. Div返回p/k；
 5. Eq判断p是否等于q；
 6. In判断p是否在r中；
 7. Mul方法返回p*k；
 8. String方法返回Point的字符串表达；
 9. Sub方法返回p-q。
*************************************************************************************/
package main

import (
	"fmt"
	"image"
)

func main() {
	pt1 := image.Pt(100, 100)
	pt2 := image.Pt(200, 200)
	pt100 := image.Pt(100, 100)
	fmt.Println(pt1.Eq(pt100), pt1.Eq(pt2))
	pt3 := pt1.Add(pt2)
	fmt.Println(pt3)
	pt4 := pt2.Div(4)
	fmt.Println(pt4)
	r := image.Rectangle{image.ZP, pt2}
	fmt.Println(pt1.In(r), pt2.In(r), pt3.In(r), pt4.In(r))
	pt5 := pt1.Mod(r)
	fmt.Println(pt5)
	pt6 := pt5.Mul(2)
	fmt.Println(pt6)
	pt7 := pt6.Sub(pt5)
	fmt.Println(pt7)
	fmt.Println(pt6.String(), pt7.String())
}
