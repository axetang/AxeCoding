/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.WriteString()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func WriteString(w Writer, s string) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 WriteString writes the contents of the string s to w, which accepts a slice of bytes.
 If w implements a WriteString method, it is invoked directly. Otherwise, w.Write is
 called exactly once.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. WriterString函数将字符串s写入w，返回值n表示写入字节数；
 2. w实现了Writer接口，即实现了Write方法；
 3. WriteString函数调用w的Write方法来完成WriteString功能。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type TestWriter struct {
	Str string
}

func WriteString(w *TestWriter, s string) (n int, err error) {
	fmt.Println("\nthis is WriteString()")
	w.Str = s
	return len(s), nil
}
func (tw *TestWriter) Write(p []byte) (n int, err error) {
	fmt.Println("\nthis is TestWriter.Write()")
	tw.Str = string(p)
	return len(p), errors.New("nil by Axe")
}

func main() {
	var tw1, tw2 TestWriter
	n, err := io.WriteString(os.Stdout, "Hello World\n")
	fmt.Println(n, "bytes writed, error is ", err)

	n, err = io.WriteString(&tw1, "Hello World")
	fmt.Println("tw1.Str: ", tw1.Str)
	fmt.Println(n, "bytes writed, error is ", err)
	n, err = io.WriteString(&tw1, "Hi, Axe")
	fmt.Println("tw1.Str: ", tw1.Str)
	fmt.Println(n, "bytes writed, error is ", err)

	WriteString(&tw2, "Hello World\n")
	fmt.Println("tw2.Str: ", tw2.Str)
}
