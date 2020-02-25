/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Seeker
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 Seeker is the interface that wraps the basic Seek method.
 Seek sets the offset for the next Read or Write to offset, interpreted according
 to whence: SeekStart means relative to the start of the file, SeekCurrent means
 relative to the current offset, and SeekEnd means relative to the end.
 Seek returns the new offset relative to the start of the file and an error, if any.
 Seeking to an offset before the start of the file is an error. Seeking to any
 positive offset is legal, but the behavior of subsequent I/O operations on the
 underlying object is implementation-dependent.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Seeker接口包装了Seek方法；
 2. Seek方法为下一次读写操作设置起始位置whence的偏移量offset；
 3. whence有三个值可选，分别是SeekStart、SeekEnd、SeekCurrent；
 4. offset依据和whence的相对位置设定为正值（向后）、负值（向前）、零值（当前位置）。
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
