/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.EqualFold
**Type: func
------------------------------------------------------------------------------------
**Definition:
func EqualFold(s, t []byte) bool
------------------------------------------------------------------------------------
**Description:
EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under
Unicode case-folding.
------------------------------------------------------------------------------------
**要点总结:
EqualFold把s和t解释成UTF-8字符串，并进行**大小写不敏感**的比较
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abc")))
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abd")))
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("aBc")))
	fmt.Println(bytes.EqualFold([]byte("a飞哥bc"), []byte("a飞哥Bc")))
}
