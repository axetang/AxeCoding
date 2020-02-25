/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Uniform
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Uniform struct {
     C color.Color
 }
 func (c *Uniform) At(x, y int) color.Color
 func (c *Uniform) Bounds() Rectangle
 func (c *Uniform) ColorModel() color.Model
 func (c *Uniform) Convert(color.Color) color.Color
 func (c *Uniform) Opaque() bool
 func (c *Uniform) RGBA() (r, g, b, a uint32)
 ------------------------------------------------------------------------------------
 **Description:
 Uniform is an infinite-sized Image of uniform color. It implements the color.Color,
 color.Model, and Image interfaces.
 Opaque scans the entire image and reports whether it is fully opaque.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Uniform类型代表一块面积无限大的具有同一色彩的图像。它实现了color.Color、color.Model和
 Image等接口；
 1.
*************************************************************************************/
package main

func main() {
}
