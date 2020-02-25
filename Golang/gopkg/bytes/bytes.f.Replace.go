/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Replace
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Replace(s, old, new []byte, n int) []byte
------------------------------------------------------------------------------------
**Description:
Replace returns a copy of the slice s with the first n non-overlapping instances of
old replaced by new. If old is empty, it matches at the beginning of the slice and
after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune slice.
If n < 0, there is no limit on the number of replacements.
------------------------------------------------------------------------------------
**要点总结:
1. Replace返回字节切片s的一个副本，并把前n个不重叠的old替换为new；如果n<0则不限制替换的数量。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("1234,234,234")
	d := bytes.Replace(s, []byte("234"), []byte("5678"), 2)
	fmt.Println(string(d))
	e := bytes.Replace(s, []byte("23"), []byte("飞哥"), -1)
	fmt.Println(string(e))
}
