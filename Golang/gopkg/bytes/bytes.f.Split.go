/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Split
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Split(s, sep []byte) [][]byte
------------------------------------------------------------------------------------
**Description:
Split slices s into all subslices separated by sep and returns a slice of the subslices
between those separators. If sep is empty, Split splits after each UTF-8 sequence. It
is equivalent to SplitN with a count of -1.
------------------------------------------------------------------------------------
**要点总结:
1. Split把字节切片s用字节切片sep分割成多个子字节切片返回。如果sep为空，Split则把s切分成每个字节
切片对应一个UTF-8字符；
2. Split等效于参数n为-1的SplitN函数。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("你,.好.,吗axe")
	for i, c := range bytes.Split(s, []byte{','}) {
		fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
	}

	println()
	s = []byte("你,.好,吗,.axe")
	for i, c := range bytes.Split(s, []byte{',', '.'}) {
		fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
	}

	println()
	for i, c := range bytes.Split(s, nil) {
		fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
	}
}
