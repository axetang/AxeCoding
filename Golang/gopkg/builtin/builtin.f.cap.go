/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.cap
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func cap(v Type) int
 ------------------------------------------------------------------------------------
 **Description:
 The cap built-in function returns the capacity of v, according to its type:
	Array: the number of elements in v (same as len(v)).
	Pointer to array: the number of elements in *v (same as len(v)).
	Slice: the maximum length the slice can reach when resliced; if v is nil,
	cap(v) is zero.
	Channel: the channel buffer capacity, in units of elements; if v is nil,
	cap(v) is zero.
 ------------------------------------------------------------------------------------
 **要点总结:
 cap 内建函数返回v的容量，返回值取决 v 的具体类型：
 	Array：v中元素的数量（与len(v)相同）。
 	Pointer to Array：*v中元素的数量（与len(v)相同）。
 	Slice：在重新切片时，切片能够达到的最大缓存长度；若v为nil，len(v) 即为零。
 	Channel：以信道元素数量为单位返回信道缓存容量；若v为nil，len(v) 即为零。
*************************************************************************************/
package main

func main() {
}
