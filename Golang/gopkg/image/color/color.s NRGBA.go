/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.NRGBA
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NRGBA struct {
     R, G, B, A uint8
 }
 func (c NRGBA) RGBA() (r, g, b, a uint32)
 ------------------------------------------------------------------------------------
 **Description:
 NRGBA represents a non-alpha-premultiplied 32-bit color.
 ------------------------------------------------------------------------------------
 **要点总结:
 NRGBA类型代表没有预乘alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
*************************************************************************************/
package main

func main() {
}
