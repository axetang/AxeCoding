/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Decode
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Decode(r io.Reader) (Image, string, error)
 ------------------------------------------------------------------------------------
 **Description:
 Decode decodes an image that has been encoded in a registered format. The string
 returned is the format name used during format registration. Format registration
 is typically done by an init function in the codec- specific package.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Decode函数解码并返回一个采用某种已注册格式编码的图像。字符串返回值是该格式注册时的名字,格式一般
 是在该编码格式的包的init函数中注册的;
 2.
*************************************************************************************/
package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("./_iofiles/Reunion.jpg")
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	config, format, err := image.DecodeConfig(reader)
	reader.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)

	reader, err = os.Open("./_iofiles/Reunion.jpg")
	defer reader.Close()
	m, str, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decode returned image type:", str)
	bounds := m.Bounds()
	fmt.Println("bounds.Min.Y", bounds.Min.Y)
	fmt.Println("bounds.Max.Y", bounds.Max.Y)
	fmt.Println("bounds.Min.X", bounds.Min.X)
	fmt.Println("bounds.Max.X", bounds.Max.X)

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
