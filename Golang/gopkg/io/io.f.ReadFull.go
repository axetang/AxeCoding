/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.ReadFull()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadFull(r Reader, buf []byte) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description: ReadFull reads exactly len(buf) bytes from r into buf. It returns the
 number of bytes copied and an error if fewer bytes were read. The error is EOF only
 if no bytes were read. If an EOF happens after reading some but not all the bytes,
 ReadFull returns ErrUnexpectedEOF. On return, n == len(buf) if and only if err == nil.
 If r returns an error having read at least len(buf) bytes, the error is dropped.
 ------------------------------------------------------------------------------------
 **要点总结:1. ReadFull从r读取len(buf)个字节；2. 如果r的可读取字节小于len(buf)，则返回unexpected
 EOF的error，以及已经读取到的字节数n；3. 只有当err==nil时，n=len(buf)；4. 同样要注意，Reader读取
 时位置受前次读取操作的影响，如果需要可以通过Seeker接口的Seek方法来重置读取位置。
*************************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 13)
	if n, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d bytes read\n", n)
	}
	fmt.Printf("%s\n\n", buf)

	// minimal read size bigger than io.Reader stream
	// Notice: now the reading position is 13+1, only 20 characters left.
	longBuf := make([]byte, 50)
	if n, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
		fmt.Println(n, "bytes read:", string(longBuf))
	}
}
