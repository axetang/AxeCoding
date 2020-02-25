/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.TeeReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func TeeReader(r Reader, w Writer) Reader
 ------------------------------------------------------------------------------------
 **Description:
 1. TeeReader returns a Reader that writes to w what it reads from r. All reads from
 r performed through it are matched with corresponding writes to w. There is no
 internal buffering - the write must complete before the read completes. Any error
 encountered while writing is reported as a read error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. TeaReader函数从参数r中读取后写入参数w，并将w转化为Reader返回；
 2. TeaReader函数实现中没有提供内部缓存，换句话说，r和w是同步的；
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	buf1 := make([]byte, 40)
	tee.Read(buf1)
	fmt.Println("buf1 is ", string(buf1))
	/*printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(tee)
	printall(&buf)*/
}
