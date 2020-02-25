/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.HasPrefix
**Type: func
------------------------------------------------------------------------------------
**Definition:
func HasSuffix(s, suffix []byte) bool
------------------------------------------------------------------------------------
**Description:
HasSuffix tests whether the byte slice s ends with suffix.
------------------------------------------------------------------------------------
**要点总结:
- HasSuffix检查s是否以suffix为后缀
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("1234567")
	fmt.Println(bytes.HasSuffix(s, []byte("567")))
	fmt.Println(bytes.HasSuffix(s, []byte("456")))
	fmt.Println(bytes.HasSuffix(s, []byte("678")))
}
