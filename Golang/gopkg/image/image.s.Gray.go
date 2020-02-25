/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Gray
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Gray struct {
     // Pix holds the image's pixels, as gray values. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *Gray) At(x, y int) color.Color
 func (p *Gray) Bounds() Rectangle
 func (p *Gray) ColorModel() color.Model
 func (p *Gray) GrayAt(x, y int) color.Gray
 func (p *Gray) Opaque() bool
 func (p *Gray) PixOffset(x, y int) int
 func (p *Gray) Set(x, y int, c color.Color)
 func (p *Gray) SetGray(x, y int, c color.Gray)
 func (p *Gray) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 Gray is an in-memory image whose At method returns color.Gray values.
 Opaque scans the entire image and reports whether it is fully opaque.
 PixOffset returns the index of the first element of Pix that corresponds to the
 pixel at (x, y).
 SubImage returns an image representing the portion of the image p visible through
 r. The returned value shares pixels with the original image.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
