/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.RGBA64
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type RGBA64 struct {
     R, G, B, A uint16
 }
 func (c RGBA64) RGBA() (r, g, b, a uint32)
 ------------------------------------------------------------------------------------
 **Description:
 RGBA64 represents a 64-bit alpha-premultiplied color, having 16 bits for each of
 red, green, blue and alpha.
 An alpha-premultiplied color component C has been scaled by alpha (A), so has valid
 values 0 <= C <= A.
 ------------------------------------------------------------------------------------
 **要点总结:
 RGBA64类型代表预乘了alpha通道的64位RGB色彩，Red、Green、Blue、Alpha各16位。
*************************************************************************************/
package main

func main() {
}
