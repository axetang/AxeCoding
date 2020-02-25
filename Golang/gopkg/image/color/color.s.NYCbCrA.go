/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.NYCbCrA
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type NYCbCrA struct {
     YCbCr
     A   uint8
 }
 func (c NYCbCrA) RGBA() (uint32, uint32, uint32, uint32)
 ------------------------------------------------------------------------------------
 **Description:
 NYCbCrA represents a non-alpha-premultiplied Y'CbCr-with-alpha color, having 8 bits
 each for one luma, two chroma and one alpha component.
 ------------------------------------------------------------------------------------
 **要点总结:
 NYCbCrA代表一个non-alpha-premultiplied Y‘CbCr-with-alpha颜色。
*************************************************************************************/
package main

func main() {
}
