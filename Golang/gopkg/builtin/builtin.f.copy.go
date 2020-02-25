/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.copy
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func copy(dst, src []Type) int
 ------------------------------------------------------------------------------------
 **Description:
 The copy built-in function copies elements from a source slice into a destination
 slice. (As a special case, it also will copy bytes from a string to a slice of
 bytes.) The source and destination may overlap. Copy returns the number of
 elements copied, which will be the minimum of len(src) and len(dst).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. copy内建函数将元素从来源切片复制到目标切片中, 返回拷贝的元素数量；
 2. 特殊情况是，它也能将字节从字符串复制到字节切片中, 来源和目标可以重叠。copy返回被复制的元素数量，
 它会是len(src)和len(dst)中较小的那个。
*************************************************************************************/
package main

import "fmt"

func main() {

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	len := copy(dst, src[4:])
	fmt.Println(len, dst)
	len = copy(dst, src[0:])
	fmt.Println(len, dst)
	len = copy(dst, src)
	fmt.Println(len, dst)
}
