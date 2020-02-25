/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.NYCbCrA
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NYCbCrA struct {
     YCbCr
     A       []uint8
     AStride int
 }
 func (p *NYCbCrA) AOffset(x, y int) int
 func (p *NYCbCrA) At(x, y int) color.Color
 func (p *NYCbCrA) ColorModel() color.Model
 func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
 func (p *NYCbCrA) Opaque() bool
 func (p *NYCbCrA) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 NYCbCrA is an in-memory image of non-alpha-premultiplied Y'CbCr-with-alpha colors.
 A and AStride are analogous to the Y and YStride fields of the embedded YCbCr.
 AOffset returns the index of the first element of A that corresponds to the pixel
 at (x, y).
 Opaque scans the entire image and reports whether it is fully opaque.
 SubImage returns an image representing the portion of the image p visible through
 r. The returned value shares pixels with the original image.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
