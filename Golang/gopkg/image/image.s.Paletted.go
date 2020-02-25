/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Paletted
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Paletted struct {
     // Pix holds the image's pixels, as palette indices. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
     // Palette is the image's palette.
     Palette color.Palette
 }
 func (p *Paletted) At(x, y int) color.Color
 func (p *Paletted) Bounds() Rectangle
 func (p *Paletted) ColorIndexAt(x, y int) uint8
 func (p *Paletted) ColorModel() color.Model
 func (p *Paletted) Opaque() bool
 func (p *Paletted) PixOffset(x, y int) int
 func (p *Paletted) Set(x, y int, c color.Color)
 func (p *Paletted) SetColorIndex(x, y int, index uint8)
 func (p *Paletted) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 Paletted is an in-memory image of uint8 indices into a given palette.
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
