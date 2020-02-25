/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.CMYK
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type CMYK struct {
     C, M, Y, K uint8
 }
 func (c CMYK) RGBA() (uint32, uint32, uint32, uint32)
 ------------------------------------------------------------------------------------
 **Description:
 CMYK represents a fully opaque CMYK color, having 8 bits for each of cyan, magenta,
 yellow and black.
 It is not associated with any particular color profile.
 ------------------------------------------------------------------------------------
 **要点总结:
 CMYK代表一个完全透明的CMYK色彩，每一个色彩8位，它与任何特定颜色模式无关。
*************************************************************************************/
package main

func main() {
}
