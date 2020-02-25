/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.NewReaderSize
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReaderSize(rd io.Reader, size int) *Reader
 **Description:
 NewReaderSize returns a new Reader whose buffer has at least the specified size.
 If the argument io.Reader is already a Reader with large enough size, it returns
 the underlying Reader.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 函数NewReaderSize的实现源代码如下：
 func NewReaderSize(rd io.Reader, size int) *Reader {
    // 已经是bufio.Reader类型，且缓存大小不小于 size，则直接返回
    b, ok := rd.(*Reader)
    if ok && len(b.buf) >= size {
        return b
    }
    // 缓存大小不会小于 minReadBufferSize （16字节）
    if size < minReadBufferSize {
        size = minReadBufferSize
    }
    // 构造一个bufio.Reader实例
    return &Reader{
        buf:          make([]byte, size),
        rd:           rd,
        lastByte:     -1,
        lastRuneSize: -1,
    }
 }
 2. NewReaderSize返回了一个新的读取器，这个读取器的缓存大小至少大于指定的大小。 如果rd参数
 （io.Reader）已经是一个有足够大缓存的io.Reader，它就会返回这个参数rd读取器的指针;
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("a string to be read at all"))
	r := bufio.NewReaderSize(rb, 81)
	var buf [128]byte
	n, _ := r.Read(buf[:])
	fmt.Println(string(buf[:n]))
}
