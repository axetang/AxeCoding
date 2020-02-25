# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)

# package color
## import "image/color"

Package color implements a basic color library.

## Index
```
Variables
func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8)
func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8)
func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8)
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8)
type Alpha
func (c Alpha) RGBA() (r, g, b, a uint32)
type Alpha16
func (c Alpha16) RGBA() (r, g, b, a uint32)
type CMYK
func (c CMYK) RGBA() (uint32, uint32, uint32, uint32)
type Color
type Gray
func (c Gray) RGBA() (r, g, b, a uint32)
type Gray16
func (c Gray16) RGBA() (r, g, b, a uint32)
type Model
func ModelFunc(f func(Color) Color) Model
type NRGBA
func (c NRGBA) RGBA() (r, g, b, a uint32)
type NRGBA64
func (c NRGBA64) RGBA() (r, g, b, a uint32)
type NYCbCrA
func (c NYCbCrA) RGBA() (uint32, uint32, uint32, uint32)
type Palette
func (p Palette) Convert(c Color) Color
func (p Palette) Index(c Color) int
type RGBA
func (c RGBA) RGBA() (r, g, b, a uint32)
type RGBA64
func (c RGBA64) RGBA() (r, g, b, a uint32)
type YCbCr
func (c YCbCr) RGBA() (uint32, uint32, uint32, uint32)
```
## Package Files

color.go ycbcr.go

