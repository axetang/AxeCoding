/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Reader
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Reader interface {
    Read(p []byte) (n int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 Reader is the interface that wraps the basic Read method. Read reads up to len(p)
 bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error
 encountered. Even if Read returns n < len(p), it may use all of p as scratch space
 during the call. If some data is available but not len(p) bytes, Read conventionally
 returns what is available instead of waiting for more. When Read encounters an error
 or end-of-file condition after successfully reading n > 0 bytes, it returns the
 number of bytes read. It may return the (non-nil) error from the same call or return
 the error (and n == 0) from a subsequent call. An instance of this general case is
 that a Reader returning a non-zero number of bytes at the end of the input stream
 may return either err == EOF or err == nil. The next Read should return 0, EOF.
 Callers should always process the n > 0 bytes returned before considering the error
 err. Doing so correctly handles I/O errors that happen after reading some bytes and
 also both of the allowed EOF behaviors. Implementations of Read are discouraged from
 returning a zero byte count with a nil error, except when len(p) == 0. Callers should
 treat a return of 0 and nil as indicating that nothing happened; in particular it
 does not indicate EOF. Implementations must not retain p.
 3 functions especially for Reader:
      func LimitReader(r Reader, n int64) Reader
      func MultiReader(readers ...Reader) Reader
      func TeeReader(r Reader, w Writer) Reader
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Reader接口只有一个方法Read；
 2. 实现Read方法的类型，通过调用Read方法来实现读取操作，注意Read方法读取是有位置效应的，第二次调用
 Read方法将会从前次调用的末尾之后的字节开始；
 3. 上次调用Read方法读取完Reader实例的所有内容，继续Read会返回EOF错误和0字节读取结果。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf1 := make([]byte, 15)
	n, err := r.Read(buf1)
	if err == nil {
		fmt.Println("buf1 is ", string(buf1), "\n", n, "bytes read.")
	}

	buf2 := make([]byte, 50)
	n, err = r.Read(buf2)
	if err == nil {
		fmt.Println("buf2 is ", string(buf2), n, "bytes read.")
	}

	buf3 := make([]byte, 10)
	n, err = r.Read(buf3)
	if err == nil {
		fmt.Println("buf3 is ", string(buf3), n, "bytes read.")
	} else {
		fmt.Println(err, n)
	}
}
