/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: color
 **Element: color.Model
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Model interface {
     Convert(c Color) Color
 }
 var (
     RGBAModel    Model = ModelFunc(rgbaModel)
     RGBA64Model  Model = ModelFunc(rgba64Model)
     NRGBAModel   Model = ModelFunc(nrgbaModel)
     NRGBA64Model Model = ModelFunc(nrgba64Model)
     AlphaModel   Model = ModelFunc(alphaModel)
     Alpha16Model Model = ModelFunc(alpha16Model)
     GrayModel    Model = ModelFunc(grayModel)
     Gray16Model  Model = ModelFunc(gray16Model)
 )
 var CMYKModel Model = ModelFunc(cmykModel)
 var NYCbCrAModel Model = ModelFunc(nYCbCrAModel)
 var YCbCrModel Model = ModelFunc(yCbCrModel)
 ------------------------------------------------------------------------------------
 **Description:
 Model can convert any Color to one from its own color model. The conversion may be
 lossy. Models for the standard color types.
 CMYKModel is the Model for CMYK colors.
 NYCbCrAModel is the Model for non-alpha-premultiplied Y'CbCr-with-alpha colors.
 YCbCrModel is the Model for Y'CbCr colors.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Model接口可以将任意Color接口转换为采用自身色彩模型的Color接口。转换可能会丢失色彩信息。
 1. CMYKModel是CMYK色彩模型的Model接口；
 2. NYCbCrAModel是non-alpha-premultiplied Y'CbCr-with-alpha色彩模型的Model接口；
 3. YCbCrModel是Y'CbCr色彩的Model接口。
*************************************************************************************/
package main

func main() {
}
