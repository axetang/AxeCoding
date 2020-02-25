/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strings
 **Element: strings.NewReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReader(s string) *Reader
 ------------------------------------------------------------------------------------
 **Description:
 NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString
 but more efficient and read-only.
 ------------------------------------------------------------------------------------
 **要点总结:
 NewReader函数返回一个从s string读取内容的Reader，和bytes.NewBufferString类似但更高效，只读。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1234567890"
	r := strings.NewReader(s)
	fmt.Println("r.Len() =", r.Len())
	r.ReadByte()
	fmt.Println("r.Len() =", r.Len())

}
