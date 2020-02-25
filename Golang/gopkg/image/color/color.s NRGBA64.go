/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.NRGBA64
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NRGBA64 struct {
     R, G, B, A uint16
 }
 func (c NRGBA64) RGBA() (r, g, b, a uint32)
 ------------------------------------------------------------------------------------
 **Description:
 NRGBA64 represents a non-alpha-premultiplied 64-bit color, having 16 bits for each
 of red, green, blue and alpha.
 ------------------------------------------------------------------------------------
 **要点总结:
 NRGBA64类型代表没有预乘alpha通道的64位RGB色彩，Red、Green、Blue、Alpha各16位。
*************************************************************************************/
package main

func main() {
}
