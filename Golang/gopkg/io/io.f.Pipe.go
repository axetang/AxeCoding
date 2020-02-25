/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Pipe()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Pipe() (*PipeReader, *PipeWriter)
 ------------------------------------------------------------------------------------
 **Description:
 Pipe creates a synchronous in-memory pipe. It can be used to connect
 code expecting an io.Reader with code expecting an io.Writer. Reads and Writes on
 the pipe are matched one to one except when multiple Reads are needed to consume a
 single Write. That is, each Write to the PipeWriter blocks until it has satisfied
 one or more Reads from the PipeReader that fully consume the written data. The data
 is copied directly from the Write to the corresponding Read (or Reads); there is no
 internal buffering. It is safe to call Read and Write in parallel with each other
 or with Close. Parallel calls to Read and parallel calls to Write are also safe:
 the individual calls will be gated sequentially.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Pipe()创建并返回一对同步的*PipeReader和*PipeWriter（分别实现了io.Reader和
 io.Writer接口）；
 2.读写同步，并不存在内部buffer，因此并发安全，但也意味着读写操作不能再同一个goroutine中并存；
 3. Read和Write都是阻塞式的，写完后才能读，读完后才能写；
 4. 对PipeReader和PipeWriter进行Close()操作，引起对方操作的EOF。如果使用CloseWithError，则返回相应Error；
 5. 写数据到管道中会堵塞，直到管道读取端读完所有数据或读取端被关闭。如果读取端关闭时带有error（即
 调用 CloseWithError 关闭），该Write返回的err就是读取端传递的error；否则err为ErrClosedPipe。
 6. 从管道中读取数据会堵塞，直到管道写入端开始写入数据或写入端被关闭。如果写入端关闭时带有error
 （即调用CloseWithError关闭），该Read返回的err就是写入端传递的error；否则err为EOF。
*************************************************************************************/
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()
	fmt.Printf("io.Pipe() returns r,w, types are %T,%T\n", r, w)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Fprintf(w, "No. %d string to be read, \n", i)
		}
		w.CloseWithError(errors.New("writer closed by goroutine self."))
	}()

	time.Sleep(500 * time.Millisecond)
	buf1 := new(bytes.Buffer)
	n, err := buf1.ReadFrom(r)
	fmt.Println(buf1.String(), n, err, "\n")

	//w was closed, so buf2 reads nothing.
	buf2 := new(bytes.Buffer)
	n, err = buf2.ReadFrom(r)
	fmt.Println(buf2.String(), n, err)
	r.Close()
}
