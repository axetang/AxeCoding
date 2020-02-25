/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.Palette
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Palette []Color
 func (p Palette) Convert(c Color) Color
 func (p Palette) Index(c Color) int
 ------------------------------------------------------------------------------------
 **Description:
 Palette is a palette of colors.
 Convert returns the palette color closest to c in Euclidean R,G,B space.
 Index returns the index of the palette color closest to c in Euclidean R,G,B,A space.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Palette类型代表一个色彩的调色板;
 1. Conver方法返回调色板中与色彩c在欧几里德RGB色彩空间最接近的色彩（实现了Model接口）；
 2. Index方法返回调色板中与色彩c在欧几里德RGB色彩空间最接近的色彩的索引。
*************************************************************************************/
package main

func main() {
}
