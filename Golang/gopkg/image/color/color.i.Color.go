/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.Color
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Color interface {
     // RGBA returns the alpha-premultiplied red, green, blue and alpha values
     // for the color. Each value ranges within [0, 0xffff], but is represented
     // by a uint32 so that multiplying by a blend factor up to 0xffff will not
     // overflow.
     //
     // An alpha-premultiplied color component c has been scaled by alpha (a),
     // so has valid values 0 <= c <= a.
     RGBA() (r, g, b, a uint32)
 }
 ------------------------------------------------------------------------------------
 **Description:
 Color can convert itself to alpha-premultiplied 16-bits per channel RGBA. The
 conversion may be lossy.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 实现了Color接口的类型可以将自身转化为预乘了alpha的16位通道的RGBA，转换可能会丢失色彩信息；
 1. Color接口只封装了一个方法RGBA，
 type Color interface {
    // 方法返回预乘了alpha的红、绿、蓝色彩值和alpha通道值，范围都在[0, 0xFFFF]。
    // 返回值类型为uint32，这样当乘以接近0xFFFF的混合参数时，不会溢出。
    RGBA() (r, g, b, a uint32)
 }
*************************************************************************************/
package main

func main() {
}
