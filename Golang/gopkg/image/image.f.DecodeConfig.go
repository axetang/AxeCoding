/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.DecodeConfig
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DecodeConfig(r io.Reader) (Config, string, error)
 ------------------------------------------------------------------------------------
 **Description:
 DecodeConfig decodes the color model and dimensions of an image that has been
 encoded in a registered format. The string returned is the format name used during
 format registration. Format registration is typically done by an init function in
 the codec-specific package.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. DecodeConfig函数解码一个采用某种已注册格式编码的图像的色彩模型和尺寸信息；
 2. 字符串返回值是该格式注册时的名字,格式一般是在该编码格式的包的init函数中注册的;
*************************************************************************************/
package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("./_iofiles/Reunion.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)
}
