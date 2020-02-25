/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.NRGBA64
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NRGBA64 struct {
     // Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *NRGBA64) At(x, y int) color.Color
 func (p *NRGBA64) Bounds() Rectangle
 func (p *NRGBA64) ColorModel() color.Model
 func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
 func (p *NRGBA64) Opaque() bool
 func (p *NRGBA64) PixOffset(x, y int) int
 func (p *NRGBA64) Set(x, y int, c color.Color)
 func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
 func (p *NRGBA64) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 NRGBA64 is an in-memory image whose At method returns color.NRGBA64 values.
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
