/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.RGBA
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type RGBA struct {
     R, G, B, A uint8
 }
 func (c RGBA) RGBA() (r, g, b, a uint32)
 ------------------------------------------------------------------------------------
 **Description:
 RGBA represents a traditional 32-bit alpha-premultiplied color, having 8 bits for
 each of red, green, blue and alpha.
 An alpha-premultiplied color component C has been scaled by alpha (A), so has valid
 values 0 <= C <= A.
 ------------------------------------------------------------------------------------
 **要点总结:
 RGBA类型代表传统的预乘了alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
*************************************************************************************/
package main

func main() {
}
