/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.NRGBA
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NRGBA struct {
     // Pix holds the image's pixels, in R, G, B, A order. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *NRGBA) At(x, y int) color.Color
 func (p *NRGBA) Bounds() Rectangle
 func (p *NRGBA) ColorModel() color.Model
 func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
 func (p *NRGBA) Opaque() bool
 func (p *NRGBA) PixOffset(x, y int) int
 func (p *NRGBA) Set(x, y int, c color.Color)
 func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
 func (p *NRGBA) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 NRGBA is an in-memory image whose At method returns color.NRGBA values.
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
