/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.HasPrefix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func HasPrefix(s, prefix []byte) bool
------------------------------------------------------------------------------------
**Description:
HasPrefix tests whether the byte slice s begins with prefix.
------------------------------------------------------------------------------------
**要点总结:
- HasPrefix检查s的前缀是否为prefix
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("abcdef")
	fmt.Println(bytes.HasPrefix(s, []byte("abc")))
	fmt.Println(bytes.HasPrefix(s, []byte("bcd")))
	fmt.Println(bytes.HasPrefix(s, []byte("dabc")))
}
