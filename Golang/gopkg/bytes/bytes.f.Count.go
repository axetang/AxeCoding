/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Count
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Count(s, sep []byte) int
------------------------------------------------------------------------------------
**Description:
Count counts the number of non-overlapping instances of sep in s. If sep is an empty
slice, Count returns 1 + the number of UTF-8-encoded code points in s.
------------------------------------------------------------------------------------
**要点总结:
Count计算子字节切片sep在字节切片s中非重叠实例的数量
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("1234567777777")
	fmt.Println(bytes.Count(s, []byte("123")))
	fmt.Println(bytes.Count(s, []byte("77")))
	fmt.Println(bytes.Count(s, []byte("777")))
}
