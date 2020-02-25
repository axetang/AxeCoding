/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.PipeWriter
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type PipeWriter struct {
    // contains filtered or unexported fields
 }
 func (w *PipeWriter) Close() error
 func (w *PipeWriter) CloseWithError(err error) error
 func (w *PipeWriter) Write(data []byte) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 A PipeWriter is the write half of a pipe.
 Close closes the writer; subsequent reads from the read half of the pipe will return
 no bytes and EOF.
 CloseWithError closes the writer; subsequent reads from the read half of the pipe
 will return no bytes and the error err, or EOF if err is nil. CloseWithError always
 returns nil.
 Write implements the standard Write interface: it writes data to the pipe, blocking
 until one or more readers have consumed all the data or the read end is closed. If
 the read end is closed with an error, that err is returned as err; otherwise err is
 ErrClosedPipe.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. PipeReader和PipeWriter都是阻塞操作，Read和Write方法需要成对匹配出现；
 2. 调用Close后关闭pipe，此后的读写操作会出错。如果使用CloseWithError关闭pipe，则返回该error；
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	reader, writer := io.Pipe()
	inputData1 := []byte("0123456789")
	inputData2 := []byte("9876543210")
	outputData := make([]byte, 12)
	var closeByCWE = errors.New("Close by CloseWithError")

	go func() {
		time.Sleep(2 * time.Second)
		writer.Write(inputData1)
	}()
	go writer.Write(inputData2)
	n, err := reader.Read(outputData)
	fmt.Println("X1 outputData:", string(outputData))
	fmt.Println("return error: ", err)
	fmt.Println(n, "bytes read.\n")

	//go writer.Write(inputData2)
	n, err = reader.Read(outputData)
	fmt.Println("X2 outputData:", string(outputData))
	fmt.Println("return error: ", err)
	fmt.Println(n, "bytes read.\n")

	writer.CloseWithError(closeByCWE)
	n, err = reader.Read(outputData)
	fmt.Println(string(outputData))
	fmt.Println("return error:", err)
	fmt.Println(n, "bytes read.")

}
