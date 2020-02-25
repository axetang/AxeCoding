/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Rectangle
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Rectangle struct {
     Min, Max Point
 }
 var ZR Rectangle
 func (r Rectangle) Add(p Point) Rectangle
 func (r Rectangle) At(x, y int) color.Color
 func (r Rectangle) Bounds() Rectangle
 func (r Rectangle) Canon() Rectangle
 func (r Rectangle) ColorModel() color.Model
 func (r Rectangle) Dx() int
 func (r Rectangle) Dy() int
 func (r Rectangle) Empty() bool
 func (r Rectangle) Eq(s Rectangle) bool
 func (r Rectangle) In(s Rectangle) bool
 func (r Rectangle) Inset(n int) Rectangle
 func (r Rectangle) Intersect(s Rectangle) Rectangle
 func (r Rectangle) Overlaps(s Rectangle) bool
 func (r Rectangle) Size() Point
 func (r Rectangle) String() string
 func (r Rectangle) Sub(p Point) Rectangle
 func (r Rectangle) Union(s Rectangle) Rectangle
 ------------------------------------------------------------------------------------
 **Description:
 A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y. It is
 well-formed if Min.X <= Max.X and likewise for Y. Points are always well-formed. A
 rectangle's methods always return well-formed outputs for well-formed inputs.
 A Rectangle is also an Image whose bounds are the rectangle itself. At returns
 color.Opaque for points in the rectangle and color.Transparent otherwise.
 ZR is the zero Rectangle.
 Add returns the rectangle r translated by p.
 At implements the Image interface.
 Bounds implements the Image interface.
 Canon returns the canonical version of r. The returned rectangle has minimum and
 maximum coordinates swapped if necessary so that it is well-formed.
 ColorModel implements the Image interface.
 Dx returns r's width.
 Dy returns r's height.
 Empty reports whether the rectangle contains no points.
 Eq reports whether r and s contain the same set of points. All empty rectangles are
 considered equal.
 In reports whether every point in r is in s.
 Inset returns the rectangle r inset by n, which may be negative. If either of r's
 dimensions is less than 2*n then an empty rectangle near the center of r will be
 returned.
 Intersect returns the largest rectangle contained by both r and s. If the two
 rectangles do not overlap then the zero rectangle will be returned.
 Overlaps reports whether r and s have a non-empty intersection.
 Size returns r's width and height.
 String returns a string representation of r like "(3,4)-(6,5)".
 Sub returns the rectangle r translated by -p.
 Union returns the smallest rectangle that contains both r and s.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Rectangle代表一个矩形，它实现了Image接口。该矩形包含所有满足Min.X<=X<Max.X且Min.Y<=Y<Max.Y
 的点。如果两个字段满足Min.X <= Max.X且Min.Y <= Max.Y，就称该实例为规范格式的。矩形的方法，当输入
 是规范格式时，总是返回规范格式的输出;
 1. Add方法返回矩形按p（作为向量）平移后的新矩形;
 2. At方法实现了Image接口的At方法；
 3. Bounds方法实现了Image接口的Bunds方法；
 4. Canon方法返回矩形的规范版本（左上&右下），方法必要时会交换坐标的最大值和最小值；
 5. ColorModel方法是实现了Image接口的ColorModel方法；
 6. Dx方法返回r的宽度；
 7. Dy方法返回r的高度；
 8. Empty方法报告矩形是否为空矩形,即内部不包含点的矩形;
 9. Eq方法报告两个矩形是否相同;
 10. In方法判断r包含的所有点都在s内，则返回真,否则返回假;
 11. Inset方法返回去掉矩形四周宽度n的框的矩形，n可为负数。如果n过大将返回靠近r中心位置的空矩形;
 12. Intersect方法返回两个矩形的交集矩形（同时被r和s包含的最大矩形）；如果r和s没有重叠会返回
 Rectangle零值l;
 13. Overlaps方法判断如果r和s有非空的交集，则返回真；否则返回假。
 14. Size方法返回r的宽度w和高度h构成的点Point{w, h}。
 15. String方法返回矩形的字符串表示，格式为"(3,4)-(6,5)";
 16. Sub方法返回矩形按p（作为向量）反向平移后的新矩形l
 17. Union方法返回同时包含r和s的最小矩形。
*************************************************************************************/
package main

import (
	"fmt"
	"image"
)

func main() {
	pt1 := image.Pt(10, 10)
	pt2 := image.Pt(100, 100)
	rect1 := image.Rectangle{pt1, pt2}
	rect2 := image.Rect(600, 600, 300, 300)
	rect3 := image.Rect(0, 0, 0, 0)
	rect4 := image.Rect(10, 10, 100, 100)
	rect5 := image.Rect(20, 20, 50, 50)
	fmt.Println(rect1, rect2)
	fmt.Println(rect1.Canon(), rect2.Canon())
	fmt.Println(rect1.Dx(), rect1.Dy(), rect2.Dx(), rect2.Dy())
	p := rect1.Size()
	fmt.Println(p)
	fmt.Println(rect1.Empty(), rect2.Empty(), rect3.Empty())
	fmt.Println(rect1.Eq(rect2), rect1.Eq(rect4))
	fmt.Println(rect1.Overlaps(rect4), rect1.Overlaps(rect2))
	fmt.Println(rect5.In(rect1), rect1.In(rect5))

	rect6 := rect1.Add(pt1)
	fmt.Println(rect6, rect2.Add(pt1))
	fmt.Println(rect1.Sub(pt1), rect2.Sub(pt1))

	fmt.Println(rect4.Intersect(rect5), rect3.Intersect(rect5))
	fmt.Println(rect1.Union(rect2), rect1)
	fmt.Println(rect1.Inset(10), rect1.Inset(-10))
	fmt.Println(rect1.String())

}
