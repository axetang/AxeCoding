/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Compare
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Compare(a, b []byte) int
------------------------------------------------------------------------------------
**Description:
Compare returns an integer comparing two byte slices lexicographically. The result
will be 0 if a==b, -1 if a < b, and +1 if a > b. A nil argument is equivalent to an
empty slice.
------------------------------------------------------------------------------------
**要点总结:
1. Compare根据字节的值比较两个字节切片的大;
2. 不同长度字节切片大小如何比较？查阅源代码，具体如下
func Compare(a, b []byte) int {
	return bytealg.Compare(a, b)
}
调用了bytealg标准包的Compare函数，具体实现如下。
func Compare(a, b []byte) int {
	l := len(a)
	//如果a比b字节数长，则让a的字节数等于b的字节数
	if len(b) < l {
		l = len(b)
	}
	//如果a的长度为0且a和b的第一个字节地址一致，说明他们是同一个字节，跳转到samebytes
	if l == 0 || &a[0] == &b[0] {
		goto samebytes
	}
	//此时要么a和b的字节数一样长，要么b的字节数比a长（len(b)>l=len(a))
	//下面的循环按照len(a)长度l来逐个字节比较大小
	for i := 0; i < l; i++ {
		c1, c2 := a[i], b[i]
		//如果a的当前字节小于b的当前字节，则返回-1，表示a<b
		if c1 < c2 {
			return -1
		}
		//反之，返回1，表示a>b
		if c1 > c2 {
			return +1
		}
	}
	//程序执行到这里，表示双方在l即len(a)长度上的比较是完全相同的
samebytes:
	//如果a的字节长度小于b的字节长度，则a<b，返回-1
	if len(a) < len(b) {
		return -1
	}
	//如果a的字节长度大于b的字节长度，则a>b，返回1
	if len(a) > len(b) {
		return +1
	}

	//如果程序能够执行到这里说明比较的a，b两个字节切片长度相同且每个字节的值完全相同即a==b，返回0
	return 0
}
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Compare([]byte{1}, []byte{2}))
	fmt.Println(bytes.Compare([]byte{1}, []byte{1}))
	fmt.Println(bytes.Compare([]byte{1}, []byte{0}))
	fmt.Println(bytes.Compare([]byte{1}, []byte{1, 2}))
	fmt.Println(bytes.Compare([]byte{2}, []byte{1, 2}))
	fmt.Println(bytes.Compare([]byte{2}, []byte{2, 2}))
	fmt.Println(bytes.Compare([]byte{2}, []byte{2, 0}))
	fmt.Println(bytes.Compare([]byte{2, 0}, []byte{2}))
}
