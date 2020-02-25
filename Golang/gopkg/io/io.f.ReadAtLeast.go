/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.ReadAtLeast
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ReadAtLeast reads from r into buf until it has read at least min bytes.
 It returns the number of bytes copied and an error if fewer bytes were read. The error
 is EOF only if no bytes were read. If an EOF happens after reading fewer than min
 bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the length of buf,
 ReadAtLeast returns ErrShortBuffer. On return, n >= min if and only if err == nil. If
 r returns an error having read at least min bytes, the error is dropped.
 ------------------------------------------------------------------------------------
 **要点总结:1. 顾名思义，从r中读取至少min个字节；2. 如果min大于buf的len，则返回shortbuffer的error；
 3. 如果min大于r的len，则返回EOF错误，但依然返回已经读取的字节数，并写入buf；4. 注意，ReadAtLeast
 函数如果连续针对一个Reader进行调用，要注意使用Seek方法重置读取起始位置，否则一旦某一次调用已经返回EOF，
 则后续的读取都是EOF，且读不到任何数据;5. 本例也展示了io包中两个常量的使用方法。
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
