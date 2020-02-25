# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package image
## import "image"

Package image implements a basic 2-D image library.

The fundamental interface is called Image. An Image contains colors, which are described in the image/color package.

Values of the Image interface are created either by calling functions such as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing image data in a format such as GIF, JPEG or PNG. Decoding any particular image format requires the prior registration of a decoder function. Registration is typically automatic as a side effect of initializing that format's package so that, to decode a PNG image, it suffices to have

import _ "image/png"
in a program's main package. The _ means to import a package purely for its initialization side effects.

See "The Go image package" for more details: https://golang.org/doc/articles/image_package.html

## Index
```
Variables
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
type Alpha
func NewAlpha(r Rectangle) *Alpha
func (p *Alpha) AlphaAt(x, y int) color.Alpha
func (p *Alpha) At(x, y int) color.Color
func (p *Alpha) Bounds() Rectangle
func (p *Alpha) ColorModel() color.Model
func (p *Alpha) Opaque() bool
func (p *Alpha) PixOffset(x, y int) int
func (p *Alpha) Set(x, y int, c color.Color)
func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
func (p *Alpha) SubImage(r Rectangle) Image
type Alpha16
func NewAlpha16(r Rectangle) *Alpha16
func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
func (p *Alpha16) At(x, y int) color.Color
func (p *Alpha16) Bounds() Rectangle
func (p *Alpha16) ColorModel() color.Model
func (p *Alpha16) Opaque() bool
func (p *Alpha16) PixOffset(x, y int) int
func (p *Alpha16) Set(x, y int, c color.Color)
func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
func (p *Alpha16) SubImage(r Rectangle) Image
type CMYK
func NewCMYK(r Rectangle) *CMYK
func (p *CMYK) At(x, y int) color.Color
func (p *CMYK) Bounds() Rectangle
func (p *CMYK) CMYKAt(x, y int) color.CMYK
func (p *CMYK) ColorModel() color.Model
func (p *CMYK) Opaque() bool
func (p *CMYK) PixOffset(x, y int) int
func (p *CMYK) Set(x, y int, c color.Color)
func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
func (p *CMYK) SubImage(r Rectangle) Image
type Config
func DecodeConfig(r io.Reader) (Config, string, error)
type Gray
func NewGray(r Rectangle) *Gray
func (p *Gray) At(x, y int) color.Color
func (p *Gray) Bounds() Rectangle
func (p *Gray) ColorModel() color.Model
func (p *Gray) GrayAt(x, y int) color.Gray
func (p *Gray) Opaque() bool
func (p *Gray) PixOffset(x, y int) int
func (p *Gray) Set(x, y int, c color.Color)
func (p *Gray) SetGray(x, y int, c color.Gray)
func (p *Gray) SubImage(r Rectangle) Image
type Gray16
func NewGray16(r Rectangle) *Gray16
func (p *Gray16) At(x, y int) color.Color
func (p *Gray16) Bounds() Rectangle
func (p *Gray16) ColorModel() color.Model
func (p *Gray16) Gray16At(x, y int) color.Gray16
func (p *Gray16) Opaque() bool
func (p *Gray16) PixOffset(x, y int) int
func (p *Gray16) Set(x, y int, c color.Color)
func (p *Gray16) SetGray16(x, y int, c color.Gray16)
func (p *Gray16) SubImage(r Rectangle) Image
type Image
func Decode(r io.Reader) (Image, string, error)
type NRGBA
func NewNRGBA(r Rectangle) *NRGBA
func (p *NRGBA) At(x, y int) color.Color
func (p *NRGBA) Bounds() Rectangle
func (p *NRGBA) ColorModel() color.Model
func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
func (p *NRGBA) Opaque() bool
func (p *NRGBA) PixOffset(x, y int) int
func (p *NRGBA) Set(x, y int, c color.Color)
func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
func (p *NRGBA) SubImage(r Rectangle) Image
type NRGBA64
func NewNRGBA64(r Rectangle) *NRGBA64
func (p *NRGBA64) At(x, y int) color.Color
func (p *NRGBA64) Bounds() Rectangle
func (p *NRGBA64) ColorModel() color.Model
func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
func (p *NRGBA64) Opaque() bool
func (p *NRGBA64) PixOffset(x, y int) int
func (p *NRGBA64) Set(x, y int, c color.Color)
func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
func (p *NRGBA64) SubImage(r Rectangle) Image
type NYCbCrA
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
func (p *NYCbCrA) AOffset(x, y int) int
func (p *NYCbCrA) At(x, y int) color.Color
func (p *NYCbCrA) ColorModel() color.Model
func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
func (p *NYCbCrA) Opaque() bool
func (p *NYCbCrA) SubImage(r Rectangle) Image
type Paletted
func NewPaletted(r Rectangle, p color.Palette) *Paletted
func (p *Paletted) At(x, y int) color.Color
func (p *Paletted) Bounds() Rectangle
func (p *Paletted) ColorIndexAt(x, y int) uint8
func (p *Paletted) ColorModel() color.Model
func (p *Paletted) Opaque() bool
func (p *Paletted) PixOffset(x, y int) int
func (p *Paletted) Set(x, y int, c color.Color)
func (p *Paletted) SetColorIndex(x, y int, index uint8)
func (p *Paletted) SubImage(r Rectangle) Image
type PalettedImage
type Point
func Pt(X, Y int) Point
func (p Point) Add(q Point) Point
func (p Point) Div(k int) Point
func (p Point) Eq(q Point) bool
func (p Point) In(r Rectangle) bool
func (p Point) Mod(r Rectangle) Point
func (p Point) Mul(k int) Point
func (p Point) String() string
func (p Point) Sub(q Point) Point
type RGBA
func NewRGBA(r Rectangle) *RGBA
func (p *RGBA) At(x, y int) color.Color
func (p *RGBA) Bounds() Rectangle
func (p *RGBA) ColorModel() color.Model
func (p *RGBA) Opaque() bool
func (p *RGBA) PixOffset(x, y int) int
func (p *RGBA) RGBAAt(x, y int) color.RGBA
func (p *RGBA) Set(x, y int, c color.Color)
func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
func (p *RGBA) SubImage(r Rectangle) Image
type RGBA64
func NewRGBA64(r Rectangle) *RGBA64
func (p *RGBA64) At(x, y int) color.Color
func (p *RGBA64) Bounds() Rectangle
func (p *RGBA64) ColorModel() color.Model
func (p *RGBA64) Opaque() bool
func (p *RGBA64) PixOffset(x, y int) int
func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
func (p *RGBA64) Set(x, y int, c color.Color)
func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
func (p *RGBA64) SubImage(r Rectangle) Image
type Rectangle
func Rect(x0, y0, x1, y1 int) Rectangle
func (r Rectangle) Add(p Point) Rectangle
func (r Rectangle) At(x, y int) color.Color
func (r Rectangle) Bounds() Rectangle
func (r Rectangle) Canon() Rectangle
func (r Rectangle) ColorModel() color.Model
func (r Rectangle) Dx() int
func (r Rectangle) Dy() int
func (r Rectangle) Empty() bool
func (r Rectangle) Eq(s Rectangle) bool
func (r Rectangle) In(s Rectangle) bool
func (r Rectangle) Inset(n int) Rectangle
func (r Rectangle) Intersect(s Rectangle) Rectangle
func (r Rectangle) Overlaps(s Rectangle) bool
func (r Rectangle) Size() Point
func (r Rectangle) String() string
func (r Rectangle) Sub(p Point) Rectangle
func (r Rectangle) Union(s Rectangle) Rectangle
type Uniform
func NewUniform(c color.Color) *Uniform
func (c *Uniform) At(x, y int) color.Color
func (c *Uniform) Bounds() Rectangle
func (c *Uniform) ColorModel() color.Model
func (c *Uniform) Convert(color.Color) color.Color
func (c *Uniform) Opaque() bool
func (c *Uniform) RGBA() (r, g, b, a uint32)
type YCbCr
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
func (p *YCbCr) At(x, y int) color.Color
func (p *YCbCr) Bounds() Rectangle
func (p *YCbCr) COffset(x, y int) int
func (p *YCbCr) ColorModel() color.Model
func (p *YCbCr) Opaque() bool
func (p *YCbCr) SubImage(r Rectangle) Image
func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
func (p *YCbCr) YOffset(x, y int) int
type YCbCrSubsampleRatio
func (s YCbCrSubsampleRatio) String() string
```
## Package Files
- format.go 
- geom.go 
- image.go 
- names.go 
- ycbcr.go

