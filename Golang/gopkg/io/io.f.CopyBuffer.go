/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.CopyBuffer()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 CopyBuffer is identical to Copy except that it stages through the
 provided buffer (if one is required) rather than allocating a temporary one. If buf
 is nil, one is allocated; otherwise if it has zero length, CopyBuffer panics.
 ------------------------------------------------------------------------------------
 **要点总结:
 1.buf是一个字符切片；
 2. buf如果只声明不初始化分配内存(即为nil)，则会自动分配一个非零len的临时buffer（需要去标准库中看下源码来进一步确认）；
 3. buf如果声明并初始化为一个0长度的字符切片，则CopyBuffer函数会抛出panic；
 4. strings.NewReader()返回一个*strings.Reader类型。
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
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	fmt.Printf("strings.NewReader() return type is %T\n", r1)
	//1. buf := make([]byte, 8)
	var buf []byte
	//3. buf := make([]byte, 0)
	//4. buf := make([]byte, 1)
	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buffer len is %d, cap is %d\n", len(buf), cap(buf))

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buffer len is %d, cap is %d\n", len(buf), cap(buf))

	srcFile, _ := os.Open("io.CopyBuffer.go")
	dstFile, _ := os.Create("io.Copy.dst.txt")
	written, err := io.CopyBuffer(dstFile, srcFile, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CopyBuffer succeed, total", written, "bytes!")
}
