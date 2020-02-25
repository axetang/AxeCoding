/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Alpha16
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Alpha16 struct {
     // Pix holds the image's pixels, as alpha values in big-endian format. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
 func (p *Alpha16) At(x, y int) color.Color
 func (p *Alpha16) Bounds() Rectangle
 func (p *Alpha16) ColorModel() color.Model
 func (p *Alpha16) Opaque() bool
 func (p *Alpha16) PixOffset(x, y int) int
 func (p *Alpha16) Set(x, y int, c color.Color)
 func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
 func (p *Alpha16) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 Alpha16 is an in-memory image whose At method returns color.Alpha16 values.
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
