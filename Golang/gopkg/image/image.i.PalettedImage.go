/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.PalettedImage
 **Type: Interface
 ------------------------------------------------------------------------------------
 **Definition:
 type PalettedImage interface {
     // ColorIndexAt returns the palette index of the pixel at (x, y).
     ColorIndexAt(x, y int) uint8
     Image
 }
 ------------------------------------------------------------------------------------
 **Description:
 PalettedImage is an image whose colors may come from a limited palette. If m is a
 PalettedImage and m.ColorModel() returns a color.Palette p, then m.At(x, y) should
 be equivalent to p[m.ColorIndexAt(x, y)]. If m's color model is not a color.Palette,
 then ColorIndexAt's behavior is undefined.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. PalettedImage接口代表一幅图像，它的像素可能来自一个有限的调色板；
 2. 如果有对象m满足PalettedImage接口，且m.ColorModel()返回的color.Model接口底层为一个Palette
 类型值（记为p），则m.At(x, y)返回值应等于p[m.ColorIndexAt(x, y)]；
 3. 如果m的色彩模型不是Palette，则ColorIndexAt的行为是不确定的。
*************************************************************************************/
package main

func main() {
}
