/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.LimitReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LimitReader(r Reader, n int64) Reader
 ------------------------------------------------------------------------------------
 **Description:
 1. LimitReader returns a Reader that reads from r but stops with EOF after n bytes.
 The underlying implementation is a *LimitedReader.
 ------------------------------------------------------------------------------------
 **要点总结:
 1.
*************************************************************************************/
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 20)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
	println()
}
