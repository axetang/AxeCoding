/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.YCbCr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type YCbCr struct {
     Y, Cb, Cr uint8
 }
 func (c YCbCr) RGBA() (uint32, uint32, uint32, uint32)
 ------------------------------------------------------------------------------------
 **Description:
 YCbCr represents a fully opaque 24-bit Y'CbCr color, having 8 bits each for one luma
 and two chroma components.
 JPEG, VP8, the MPEG family and other codecs use this color model. Such codecs often
 use the terms YUV and Y'CbCr interchangeably, but strictly speaking, the term YUV
 applies only to analog video signals, and Y' (luma) is Y (luminance) after applying
 gamma correction.
 Conversion between RGB and Y'CbCr is lossy and there are multiple, slightly
 different formulae for converting between the two. This package follows the JFIF
 specification at https://www.w3.org/Graphics/JPEG/jfif3.pdf.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. YcbCr代表完全不透明的24位Y'CbCr色彩；每个色彩都有1个亮度成分和2个色度成分，分别用8位字节表示;
 2. JPEG、VP8、MPEG家族和其他编码方式使用本色彩模型。这些编码通常将Y'CbCr 和YUV两个色彩模型等同
 使用（Y=Y'=黄、U=Cb=青、V=Cr=品红）。但严格来说，YUV模只用于模拟视频信号，Y'是经过伽玛校正的Y。
 RGB和Y'CbCr色彩模型之间的转换会丢失色彩信息。两个色彩模型之间的转换有多个存在细微区别的算法。本包
 采用JFIF算法，参见http://www.w3.org/Graphics/JPEG/jfif3.pdf.
*************************************************************************************/
package main

func main() {
}
