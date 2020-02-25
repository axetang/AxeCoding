/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.SplitAfter
**Type: func
------------------------------------------------------------------------------------
**Definition:
func SplitAfter(s, sep []byte) [][]byte
------------------------------------------------------------------------------------
**Description:
SplitAfter slices s into all subslices after each instance of sep and returns a slice
of those subslices. If sep is empty, SplitAfter splits after each UTF-8 sequence. It
is equivalent to SplitAfterN with a count of -1.
------------------------------------------------------------------------------------
**要点总结:
1. SplitAfter用sep作为后缀把s切分成多个字节切片返回;
2. 如果sep为空，则把s切分成每个字节切片对应一个UTF-8字符。
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
	s := []byte("1,2,3")
	sep := []byte{','}
	print(bytes.SplitAfter(s, sep))
	println()
	print(bytes.SplitAfter(s, nil))
}
