/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.RegisterFormat
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
 ------------------------------------------------------------------------------------
 **Description:
 RegisterFormat registers an image format for use by Decode. Name is the name of the
 format, like "jpeg" or "png". Magic is the magic prefix that identifies the format's
 encoding. The magic string can contain "?" wildcards that each match any one byte.
 Decode is the function that decodes the encoded image. DecodeConfig is the function
 that decodes just its configuration.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RegisterFormat函数注册一个供Decode函数使用的图片格式；
 2. 参数name是格式的名字，如"jpeg"或"png"，参数magic是该格式编码的魔术前缀，该字符串可以包含"?"
 通配符，每个通配符匹配一个字节；
 3. decode函数用于解码图片；decodeConfig函数只解码图片的配置。
*************************************************************************************/
package main

func main() {
}
