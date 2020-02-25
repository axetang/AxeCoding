/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.len
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func len(v Type) int
 ------------------------------------------------------------------------------------
 **Description:
 The len built-in function returns the length of v, according to its type:
	Pointer to array: the number of elements in *v (even if v is nil).
	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
	String: the number of bytes in v.
	Channel: the number of elements queued (unread) in the channel buffer;
	if v is nil, len(v) is zero.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. len内建函数返回v的长度，这取决于具体类型：
 	数组：v中元素的数量。
	数组指针：*v中元素的数量（即使v为 nil）。
	切片或映射：v中元素的数量；若v为 nil，len(v) 即为零。
	字符串：v中字节的数量。
	信道：信道缓存中队列（未读取）元素的数量；若v为 nil，len(v) 即为零。
*************************************************************************************/
package main

func main() {
}
