/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package:
 **Element:
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:

 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:1.
*************************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	str := "some io.Reader stream to be read\n"
	r := strings.NewReader(str)
	println("str's len is : ", len(str))

	buf := make([]byte, 33)
	if n, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d bytes read\n", n)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	r.Seek(10, io.SeekStart)
	shortBuf := make([]byte, 10)
	if n, err := io.ReadAtLeast(r, shortBuf, 20); err != nil {
		fmt.Printf("error: type=%T, content=%s\n", err, err)
		fmt.Print(string(shortBuf))
		fmt.Println(shortBuf, "\n", n, "bytes read")
	}
	println()

	// minimal read size bigger than io.Reader stream
	r.Seek(-28, io.SeekEnd)
	longBuf := make([]byte, 40)
	if n, err := io.ReadAtLeast(r, longBuf, 40); err != nil {
		fmt.Println("error::", err)
		fmt.Print(string(longBuf))
		fmt.Println(longBuf, "\n", n, "bytes read")
	}
}
