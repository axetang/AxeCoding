/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.RangeTable
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type RangeTable struct {
     R16         []Range16
     R32         []Range32
     LatinOffset int // number of entries in R16 with Hi <= MaxLatin1
 }
 ------------------------------------------------------------------------------------
 **Description:
 RangeTable defines a set of Unicode code points by listing the ranges of code points
 within the set. The ranges are listed in two slices to save space: a slice of 16-bit
 ranges and a slice of 32-bit ranges. The two slices must be in sorted order and
 non-overlapping. Also, R32 should contain only values >= 0x10000 (1<<16).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RangeTable定义一个Unicode码点集合，包含16位和32位两个范围列表。这两个列表必须经过排序而且
 不能重叠。R32中只能包含大于16位的值；
 2. 后续刷新。
*************************************************************************************/
package main

func main() {
}
