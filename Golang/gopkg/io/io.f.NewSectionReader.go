/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.NewSectionReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
 ------------------------------------------------------------------------------------
 **Description:
 NewSectionReader returns a SectionReader that reads from r starting at offset off
 and stops with EOF after n bytes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewSectionReader返回一个SectionReader，它从r中的偏移量off处读取n个字节后以EOF停止。

*************************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 3, 50)

	n, err := io.Copy(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(n, "bytes copied.")
	}
	println()

	s1 := io.NewSectionReader(r, 40, 50)
	n, err = io.Copy(os.Stdout, s1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("error == nil", n, "bytes copied.")
	}
	println()
}
