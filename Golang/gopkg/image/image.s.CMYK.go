/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.CMYK
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type CMYK struct {
     // Pix holds the image's pixels, in C, M, Y, K order. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *CMYK) At(x, y int) color.Color
 func (p *CMYK) Bounds() Rectangle
 func (p *CMYK) CMYKAt(x, y int) color.CMYK
 func (p *CMYK) ColorModel() color.Model
 func (p *CMYK) Opaque() bool
 func (p *CMYK) PixOffset(x, y int) int
 func (p *CMYK) Set(x, y int, c color.Color)
 func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
 func (p *CMYK) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 CMYK is an in-memory image whose At method returns color.CMYK values.
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
