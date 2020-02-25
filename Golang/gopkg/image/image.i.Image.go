/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Image
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Image interface {
     // ColorModel returns the Image's color model.
     ColorModel() color.Model
     // Bounds returns the domain for which At can return non-zero color.
     // The bounds do not necessarily contain the point (0, 0).
     Bounds() Rectangle
     // At returns the color of the pixel at (x, y).
     // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
     // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
     At(x, y int) color.Color
 }
 ------------------------------------------------------------------------------------
 **Description:
 Image is a finite rectangular grid of color.Color values taken from a color model.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Image接口定义说明如下：
 type Image interface {
    // ColorModel方法返回图像的色彩模型
    ColorModel() color.Model
    // Bounds方法返回图像的范围，范围不一定包括点(0, 0)
    Bounds() Rectangle
    // At方法返回(x, y)位置的色彩
    // At(Bounds().Min.X, Bounds().Min.Y)返回网格左上角像素的色彩
    // At(Bounds().Max.X-1, Bounds().Max.Y-1) 返回网格右下角像素的色彩
    At(x, y int) color.Color
 }
 2. Image接口封装了三个方法：ColorModel，Bounds, At.
*************************************************************************************/
package main

func main() {
}
