/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.WriterTo
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 WriterTo is the interface that wraps the WriteTo method.
 WriteTo writes data to w until there's no more data to write or when an error
 occurs. The return value n is the number of bytes written. Any error encountered
 during the write is also returned.
 The Copy function uses WriterTo if available.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. WriterTo接口包装了WriteTo方法;
 2. WriteTo将数据写入w中，直到没有数据可读或发生错误。其返回值n为写入的字节数。 在写入过程中遇到
 的任何错误也将被返回;
 3. 如果Reader一方的WriterTo方法可用，Copy函数就会使用它而不是Writer一方的Write方法.在标准库
 源代码中可以看到Copy函数调用CopyBuffer函数，CopyBuffer函数判断Reader一方是否支持WriteTo方法，
 如果有可用的WriteTo方法，就直接调用。如果没有就调用Writer一方的Write方法来完成Write操作;
 4. 根据标准库源码，首先做Reader一方的WriteTo方法判断，如果可用就直接使用WriteTo方法（而不是Writer
 方的Write方法）完成后续操作；第二步才会判断Writer一方是否有ReadFrom方法可用，如果有才会使用
 ReadFrom方法（而不是Reader方的Read方法）。注意，如果第一步已经判断WriteTo方法可用，就不会继续
 走到第二步了。
*************************************************************************************/
package main

import (
	"fmt"
	"io"
)

type MyWriter struct {
	W io.Writer
}
type MyReader struct {
	R io.Reader
}

//可以尝试注释掉WriteTo方法看看运行结果
/*func (mr *MyReader) WriteTo(w io.Writer) (int64, error) {
	fmt.Println("this is mr.WriteTo")
	return 1, nil
}*/

func (mr *MyReader) Read(p []byte) (int, error) {
	fmt.Println("this is mr.Read()")
	return 1, nil
}

func (mw *MyWriter) Write(p []byte) (int, error) {
	fmt.Println("this is mw.Write()")
	return 1, nil
}
func (mw *MyWriter) ReadFrom(r io.Reader) (int64, error) {
	fmt.Println("this is mw.ReadFrom")
	return 1, nil
}

func main() {
	var mw MyWriter
	var mr MyReader
	//reader := strings.NewReader("This is a reader by strings.NewReader()")
	n, err := io.Copy(&mw, &mr)
	fmt.Println(n, err)

}
