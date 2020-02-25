/*************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.CopyN()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 CopyN copies n bytes (or until an error) from src to dst. It returns
 the number of bytes copied and the earliest error encountered while copying. On
 return, written == n if and only if err == nil. If dst implements the ReaderFrom
 interface, the copy is implemented using it.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. golang的函数参数顺序特点之一是dst在前，src在后；
 2. 参数n表示拷贝n个字节；
 3. 如果err==nil，则返回值written<=n；3.如果n大于src的长度即遇到EOF,程序异常退出；
 4. CopyN也是通过调用Copy函数完成功能的：written, err = Copy(dst, LimitReader(src, n))
************************************************************************************/

package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	written, err := io.CopyN(os.Stdout, r, 10)
	if err != nil {
		log.Fatal(err)
	}
	println("\ntotal copied bytes:", written)

	written, err = io.CopyN(os.Stdout, r, 1000)
	if err != nil {
		log.Fatal(err)
	}
	println("\ntotal copied bytes:", written)
}
