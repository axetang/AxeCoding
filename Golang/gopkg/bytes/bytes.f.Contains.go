/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Contains
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Contains(b, subslice []byte) bool
------------------------------------------------------------------------------------
**Description:
Contains reports whether subslice is within b.
------------------------------------------------------------------------------------
**要点总结:
1. Contains检查字节切片b是否包含子字节切片subslice
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("12345678")
	s1 := []byte("456")
	s2 := []byte("789")
	fmt.Println(bytes.Contains(b, s1))
	fmt.Println(bytes.Contains(b, s2))
}
