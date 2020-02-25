/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.YCbCr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type YCbCr struct {
     Y, Cb, Cr      []uint8
     YStride        int
     CStride        int
     SubsampleRatio YCbCrSubsampleRatio
     Rect           Rectangle
 }
 func (p *YCbCr) At(x, y int) color.Color
 func (p *YCbCr) Bounds() Rectangle
 func (p *YCbCr) COffset(x, y int) int
 func (p *YCbCr) ColorModel() color.Model
 func (p *YCbCr) Opaque() bool
 func (p *YCbCr) SubImage(r Rectangle) Image
 func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
 func (p *YCbCr) YOffset(x, y int) int
 ------------------------------------------------------------------------------------
 **Description:
 YCbCr is an in-memory image of Y'CbCr colors. There is one Y sample per pixel, but each Cb and Cr sample can span one or more pixels. YStride is the Y slice index delta between vertically adjacent pixels. CStride is the Cb and Cr slice index delta between vertically adjacent pixels that map to separate chroma samples. It is not an absolute requirement, but YStride and len(Y) are typically multiples of 8, and:
 For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
 For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
 For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
 For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
 For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
 For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.
 COffset returns the index of the first element of Cb or Cr that corresponds to the
 pixel at (x, y).
 SubImage returns an image representing the portion of the image p visible through r.
 The returned value shares pixels with the original image.
 YOffset returns the index of the first element of Y that corresponds to the pixel
 at (x, y).
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
