/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.PipeReader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type PipeReader struct {
	// contains filtered or unexported fields
	p *pipe
 }
 func (r *PipeReader) Close() error
 func (r *PipeReader) CloseWithError(err error) error
 func (r *PipeReader) Read(data []byte) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 A PipeReader is the read half of a pipe.
 Close closes the reader; subsequent writes to the write half of the pipe will return
 the error ErrClosedPipe.
 CloseWithError closes the reader， subsequent writes to the write half of the pipe
 will return the error err.
 Read implements the standard Read interface: it reads data from the pipe, blocking
 until a writer arrives or the write end is closed. If the write end is closed with
 an error, that error is returned as err; otherwise err is EOF.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 结构体PipeReader中唯一的成员p *pipe，pipe是io包中定义的一个结构体，查询标准包中定义如下：
 type pipe struct {
	wrMu sync.Mutex // Serializes Write operations
	wrCh chan []byte
	rdCh chan int

	once sync.Once // Protects closing done
	done chan struct{}
	rerr atomicError
	werr atomicError
 }
 2. PipeReader实现了Read方法从而实现了io.Reader接口；
 3. 在PipeWriter一方调用Close()方法表示写入结束并阻塞等待Read，PipeReader通过Read方法进行阻塞式
 Read操作，会一次读完PipeWriter写入的全部内容。此后如果继续在PipeReader一方进行Read，不能读出任何
 数据，同时err为nil；
 4. 如果PipeWriter一方调用了CloseWithError方法完成写入，则此后PipeReader通过Read方法读取完所有
 写入内容后，也会发挥CloseWithError创建的error。如果此后继续读入，则不能读出任何内容，同时继续返回
 前述CloseWithError创建的error；
 5. 从pipe结构体定义来看，pipe应该也是使用channel来实现的，具体需要后续阅读标准包源代码后再确认；
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()
	fmt.Printf("io.Pipe() returns r,w, types are %T,%T\n\n", r, w)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Fprintf(w, "No. %d string to be read, \n", i)
		}
		//w.CloseWithError(errors.New("writer closed by goroutine self."))
		w.Close()
	}()

	buf1 := new(bytes.Buffer)
	n, err := buf1.ReadFrom(r)
	fmt.Println(buf1.String(), n, "bytes read, ", err, "\n")
	//fmt.Println(buf1.String(), n, "bytes read, ", err.Error(), "\n")

	//w was closed, so buf2 reads nothing.
	time.Sleep(500 * time.Millisecond)
	buf2 := new(bytes.Buffer)
	n, err = buf2.ReadFrom(r)
	fmt.Println(buf2.String(), n, "bytes read. ", err)
	//fmt.Println(buf2.String(), n, "bytes read. ", err.Error())
	//r.Close()
}
