/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.RGBA64
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type RGBA64 struct {
     // Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *RGBA64) At(x, y int) color.Color
 func (p *RGBA64) Bounds() Rectangle
 func (p *RGBA64) ColorModel() color.Model
 func (p *RGBA64) Opaque() bool
 func (p *RGBA64) PixOffset(x, y int) int
 func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
 func (p *RGBA64) Set(x, y int, c color.Color)
 func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
 func (p *RGBA64) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 RGBA64 is an in-memory image whose At method returns color.RGBA64 values.
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
