/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Repeat
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Repeat(b []byte, count int) []byte
------------------------------------------------------------------------------------
**Description:
Repeat returns a new byte slice consisting of count copies of b.
------------------------------------------------------------------------------------
**要点总结:
1. Repeat把byte切片b复制count次组合成一个新的字节切片并返回。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("123")
	fmt.Println(string(bytes.Repeat(b, 3)))
	fmt.Println(string(bytes.Repeat(b, 5)))
}
