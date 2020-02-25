/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Runes
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Runes(s []byte) []rune
------------------------------------------------------------------------------------
**Description:
Runes returns a slice of runes (Unicode code points) equivalent to s.
------------------------------------------------------------------------------------
**要点总结:
1. Runes把s解释为UTF-8编码的字节序列，并返回对应的Unicode切片。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("你好世界")
	fmt.Printf("%d, %s\n", len(s), string(s))
	r := bytes.Runes(s)
	fmt.Printf("%d, %s\n", len(r), string(r))
}
