/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Equal
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Equal(a, b []byte) bool
------------------------------------------------------------------------------------
**Description:
Equal returns a boolean reporting whether a and b are the same length and contain the
same bytes. A nil argument is equivalent to an empty slice.
------------------------------------------------------------------------------------
**要点总结:
1. 比较两个字节切片是否相等，如果参数为nil，则等同于空的字节切片
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2, 3}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, nil))
	fmt.Println(bytes.Equal([]byte{}, nil))
	fmt.Println(bytes.Equal(nil, nil))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))
}
