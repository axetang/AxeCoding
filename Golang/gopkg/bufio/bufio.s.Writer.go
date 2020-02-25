/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.Writer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Writer struct {
    // contains filtered or unexported fields
	err error
	buf []byte
	n   int
	wr  io.Writer
 }
 func NewWriter(w io.Writer) *Writer
 func NewWriterSize(w io.Writer, size int) *Writer
 func (b *Writer) Available() int
 func (b *Writer) Buffered() int
 func (b *Writer) Flush() error
 func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
 func (b *Writer) Reset(w io.Writer)
 func (b *Writer) Size() int
 func (b *Writer) Write(p []byte) (nn int, err error)
 func (b *Writer) WriteByte(c byte) error
 func (b *Writer) WriteRune(r rune) (size int, err error)
 func (b *Writer) WriteString(s string) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 Writer implements buffering for an io.Writer object. If an error occurs writing
 to a Writer, no more data will be accepted and all subsequent writes, and Flush,
 will return the error. After all data has been written, the client should call
 the Flush method to guarantee all data has been forwarded to the underlying
 io.Writer.
 NewWriter returns a new Writer whose buffer has the default size.
 NewWriterSize returns a new Writer whose buffer has at least the specified size.
 If the argument io.Writer is already a Writer with large enough size, it returns
 the underlying Writer.
 Available returns how many bytes are unused in the buffer.
 Buffered returns the number of bytes that have been written into the current buffer.
 Flush writes any buffered data to the underlying io.Writer.
 ReadFrom implements io.ReaderFrom. If the underlying writer supports the
 ReadFrom method, and b has no buffered data yet, this calls the underlying
 ReadFrom without buffering.
 Reset discards any unflushed buffered data, clears any error, and resets b to
 write its output to w.
 Size returns the size of the underlying buffer in bytes.
 Write writes the contents of p into the buffer. It returns the number of bytes
 written. If nn < len(p), it also returns an error explaining why the write is
 short.
 WriteByte writes a single byte.
 WriteRune writes a single Unicode code point, returning the number of bytes
 written and any error.
 WriteString writes a string. It returns the number of bytes written. If the
 count is less than len(s), it also returns an error explaining why the write is
 short.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Writer实现了io.Writer对象的缓存。 如果在写数据到Writer的时候出现了一个错误，不会再有数据被写
 进来了， 并且所有随后的写操作都会返回error。当所有数据被写入后，客户端应调用 Flush 方法以确保
 所有数据已转为基本的io.Writer;特别的，在bufio.go中定义的内部常量defaultBufSize=4096；
 2. NewWriter返回一个新的，有默认尺寸缓存的Writer；
 3. NewWriterSize返回一个新的Writer，它的缓存一定大于指定的size参数。 如果io.Writer参数已经
 是足够大的有缓存的Writer了，函数就会返回它底层的Writer；
 4. Available返回buffer中有多少尚未使用的字节总数；
 5. Buffered返回已经写入到当前缓存的字节数；
 6. Flush将缓存上的所有数据写入到底层的io.Writer中。
 7. ReadFrom实现了io.ReaderFrom。
 8. Write将p中的内容写入到缓存中。 它返回写入的字节数。如果n<len(p), 它会返回错误;
 9. WriteByte向缓存中写入单个字节;
 10. WriteRune向缓存中写入单个的Unicode代码，返回写入的字节数和遇到的错误;
 11. WriteString向缓存中写入一个string，返回写入的字节数。 如果字节数比len(s)少，它就会返回error.
 *************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func TestAvailableAndBuffered() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	fmt.Println("Available", "Buffered")
	fmt.Println(w.Available(), w.Buffered())
	w.WriteByte('1')
	fmt.Println(w.Available(), w.Buffered())
	w.WriteString("2345")
	fmt.Println(w.Available(), w.Buffered())
}

func TestWriteAndFlush() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	w.Write([]byte("0123456789"))
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}

func TestWriteByte() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	w.WriteByte('a')
	w.Flush()
	fmt.Println(string(wb.Bytes()))
}

func TestWriteRune() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	n, err := w.WriteRune('你')
	if err != nil {
		return
	}
	w.Flush()
	fmt.Println(n)
	fmt.Println(wb.Bytes())
}

func TestWriteString() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	n, err := w.WriteString("123456")
	if err != nil {
		return
	}
	fmt.Println(n)
	fmt.Println(wb.Bytes())
	w.Flush()
	fmt.Println(wb.Bytes())

}

func main() {
	/*fmt.Println("---------Available & Buffered-----------")
	TestAvailableAndBuffered()
	fmt.Println("---------Write & Flush-----------")
	TestWriteAndFlush()*/
	/*fmt.Println("---------Reset-----------")
	TestReset()
	fmt.Println("---------Size-----------")
	TestSize()
	fmt.Println("---------ReadFrom-----------")
	TestReadFrom()*/
	fmt.Println("---------WriteByte-----------")
	TestWriteByte()
	fmt.Println("---------WriteRune-----------")
	TestWriteRune()
	fmt.Println("---------WriteString-----------")
	TestWriteString()

}
