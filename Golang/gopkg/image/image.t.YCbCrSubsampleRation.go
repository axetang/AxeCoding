/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.YCbCrSubsampleRation
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type YCbCrSubsampleRatio int
 const (
     YCbCrSubsampleRatio444 YCbCrSubsampleRatio = iota
     YCbCrSubsampleRatio422
     YCbCrSubsampleRatio420
     YCbCrSubsampleRatio440
     YCbCrSubsampleRatio411
     YCbCrSubsampleRatio410
 )
 func (s YCbCrSubsampleRatio) String() string
 ------------------------------------------------------------------------------------
 **Description:
 YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
