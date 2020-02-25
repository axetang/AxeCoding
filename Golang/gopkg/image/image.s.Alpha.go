/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: image
 **Element: image.Alpha
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Alpha struct {
     // Pix holds the image's pixels, as alpha values. The pixel at
     // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
     Pix []uint8
     // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
     Stride int
     // Rect is the image's bounds.
     Rect Rectangle
 }
 func (p *Alpha) AlphaAt(x, y int) color.Alpha
 func (p *Alpha) At(x, y int) color.Color
 func (p *Alpha) Bounds() Rectangle
 func (p *Alpha) ColorModel() color.Model
 func (p *Alpha) Opaque() bool
 func (p *Alpha) PixOffset(x, y int) int
 func (p *Alpha) Set(x, y int, c color.Color)
 func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
 func (p *Alpha) SubImage(r Rectangle) Image
 ------------------------------------------------------------------------------------
 **Description:
 Alpha is an in-memory image whose At method returns color.Alpha values.
 Opaque scans the entire image and reports whether it is fully opaque.
 PixOffset returns the index of the first element of Pix that corresponds to the
 pixel at (x, y).
 SubImage returns an image representing the portion of the image p visible through r.
 The returned value shares pixels with the original image.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. AlphaAt(x, y int) color.Alpha      //获取指定点的透明度
 2. At(x, y int) color.Color　　　//获取指定点的color(指定点的红绿蓝的透明度)
 3. Bounds() Rectangle　　　//获取alpha的边界
 4. ColorModel() color.Model　//获取alpha的颜色模型
 5. Opaque() bool　　　　　　//检查alpha是否完全不透明
 6. PixOffset(x, y int) int　　　//获取指定像素相对于第一个像素的相对偏移量
 7. Set(x, y int, c color.Color)　//设定指定位置的color
 8. SetAlpha(x, y int, c color.Alpha)  //设定指定位置的alpha
 9. SubImage(r Rectangle) Image   //获取p图像中被r覆盖的子图像，父图像和子图像公用像素
*************************************************************************************/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

const (
	dx = 500
	dy = 200
)

func main() {

	file, err := os.Create("./_iofiles/test.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	alpha := image.NewAlpha(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			alpha.Set(x, y, color.Alpha{uint8(x % 256)}) //设定alpha图片的透明度
		}
	}
	fmt.Println(alpha.At(400, 100))    //144 在指定位置的像素
	fmt.Println(alpha.Bounds())        //(0,0)-(500,200)，图片边界
	fmt.Println(alpha.Opaque())        //false，是否图片完全透明
	fmt.Println(alpha.PixOffset(1, 1)) //501，指定点相对于第一个点的距离
	fmt.Println(alpha.Stride)          //500，两个垂直像素之间的距离
	jpeg.Encode(file, alpha, nil)      //将image信息写入文件中
}
