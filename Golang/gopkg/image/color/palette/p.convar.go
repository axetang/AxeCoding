/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: palette
 **Element: palette.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 var Plan9 = []color.Color{
	 // 256 elements not displayed
 }
 var WebSafe = []color.Color{
	 //216 elements not displayed
 }
 ------------------------------------------------------------------------------------
 **Description:
 Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4
 subdivision, with 4 shades in each subcube. Compared to the WebSafe, the idea is
 to reduce the color resolution by dicing the color cube into fewer cells, and to
 use the extra space to increase the intensity resolution. This results in 16 gray
 shades (4 gray subcubes with 4 samples in each), 13 shades of each primary and
 secondary color (3 subcubes with 4 samples plus black) and a reasonable selection
 of colors covering the rest of the color cube. The advantage is better representation
 of continuous tones.
 This palette was used in the Plan 9 Operating System, described at
 https://9p.io/magic/man2html/6/color
 WebSafe is a 216-color palette that was popularized by early versions of Netscape
 Navigator. It is also known as the Netscape Color Cube.
 See https://en.wikipedia.org/wiki/Web_colors#Web-safe_colors for details.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. palette包提供了标准的调色板;
 1. Plan9是一个256色的调色板，将24位的RGB色彩立方划分为4x4x4的子立方，每个子立方体都有4个色调;
 比起WebSafe，本方案通过将颜色立方划分成更少的方格减少了颜色的分辨率，以便用额外的空间换取增加强度的
 分辨率。结果是灰色有16个色调（4个灰色子立方每个有4个采样），初级和次级颜色各有13个色调（3个子立方每
 个有4个采样，再加上黑色），加上对色彩立方其余部分的合理色彩选择，实现了更好的连续色调的表现.本调色板
 用于Plan9操作系统，参见http://plan9.bell-labs.com/magic/man2html/6/color;
 2. WebSafe是一个216色的调色板，被早期版本的Netscape Navigator（一种浏览器）广泛使用，它也被称
 为Netcape色彩立方。细节参见http://en.wikipedia.org/wiki/Web_colors#Web-safe_colors;
 3. 具体参见源码。
*************************************************************************************/
package main

func main() {
}
