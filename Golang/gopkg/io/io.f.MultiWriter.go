/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.MultiWriter
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func MultiWriter(writers ...Writer) Writer
 ------------------------------------------------------------------------------------
 **Description:
 MultiWriter creates a writer that duplicates its writes to all the provided writers,
 similar to the Unix tee(1) command.
 Each write is written to each listed writer, one at a time. If a listed writer
 returns an error, that overall write operation stops and returns the error; it does
 not continue down the list.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. MultiWriter函数将多个writer组成一个新的writer返回；
 2. 在向这个新的writer进行Write操作时，将按顺序向组成这个新writer的所有writer依次一致全部写入；
 3. 当任何一个成员writer写入出现错误时，全部的writer操作停止并返回错误，而不是会继续写完所有成员；
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2, buf3 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2, &buf3)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
	fmt.Print(buf3.String())
}
