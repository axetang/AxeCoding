/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.SplitN
**Type: func
------------------------------------------------------------------------------------
**Definition:
func SplitN(s, sep []byte, n int) [][]byte
------------------------------------------------------------------------------------
**Description:
SplitN slices s into subslices separated by sep and returns a slice of the subslices
between those separators. If sep is empty, SplitN splits after each UTF-8 sequence.
The count determines the number of subslices to return:
	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
	n == 0: the result is nil (zero subslices)
	n < 0: all subslices
------------------------------------------------------------------------------------
**要点总结:
1. SplitN把s用sep分割成多个字节切片返回, 如果sep为nil，Split则把s切分成每个子字节切片对应一个
UTF-8字符;
2. SplitN的参数n决定返回切片的长度：
	n > 0: 最多返回n个子切片；最后可能子切片可能包含未切分的字节序列。
	n == 0: 返回空切片
	n < 0: 返回所有的子切片
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func print(s [][]byte) {
	if len(s) == 0 {
		fmt.Println("nil")
		return
	}
	for i, c := range s {
		fmt.Printf("%d: %s\n", i, string(c))
	}
}

func main() {
	s := []byte("1,2,3,4")
	sep := []byte{','}
	print(bytes.SplitN(s, sep, 2))
	println()
	print(bytes.SplitN(s, sep, 0))
	println()
	print(bytes.SplitN(s, sep, -1))
}
