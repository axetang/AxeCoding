/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strings
 **Element: strings.Reader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Reader struct {
	// contains filtered or unexported fields
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
 }
 A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo, io.ByteScanner,
 and io.RuneScanner interfaces by reading from a string.
 func NewReader(s string) *Reader
 func (r *Reader) Len() int
 func (r *Reader) Read(b []byte) (n int, err error)
 func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
 func (r *Reader) ReadByte() (byte, error)
 func (r *Reader) ReadRune() (ch rune, size int, err error)
 func (r *Reader) Reset(s string)
 func (r *Reader) Seek(offset int64, whence int) (int64, error)
 func (r *Reader) Size() int64
 func (r *Reader) UnreadByte() error
 func (r *Reader) UnreadRune() error
 func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString
 but more efficient and read-only.
 Len returns the number of bytes of the unread portion of the string.
 Reset resets the Reader to be reading from s.
 Seek implements the io.Seeker interface.
 Size returns the original length of the underlying string. Size is the number of
 bytes available for reading via ReadAt. The returned value is always the same and is
 not affected by calls to any other method.
 WriteTo implements the io.WriterTo interface.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 结构体Reader从参数s string读取内容来实现io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
 io.ByteScanner和io.RuneScanner等接口；
 1. Len方法返回结构体Reader的未读字节数；
 2. Read方法从底层字符串读取内容到参数p []byte，读取长度为len(b)；
 3. ReadAt方法从底层字符串读取位置偏移off处读取内容，读取疮毒为len(b)；
 4. ReadByte从底层字符串读取一个字节并返回；
 5. ReadRune从底层字符串读取一个rune并返回；
 6. Reset方法重置底层字符串为参数s；
 7. Seek实现了io.Seeker接口；
 8. Size方法返回结构体Reader的底层字符串初始长度；
 9. UnreadByte将结构体Reader上次读取字符串的位置回退一个字节
 10. UnreadRune将结构体Reader上次读取的UTF-8字符串位置回退一个UTF-8字符
 11. WriteTo实现了io.WriterTo接口
*************************************************************************************/
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "0123456789"
	r := strings.NewReader(s)

	fmt.Println("-------Len&ReadByte-----------")
	fmt.Println("r.Len() =", r.Len())
	r.ReadByte()
	fmt.Println("r.Len() =", r.Len())

	fmt.Println("-------Read-----------")
	s = "0123456789"
	r = strings.NewReader(s)
	var buf [3]byte
	fmt.Println("buf.Len() =", len(buf))
	n, _ := r.Read(buf[:])
	fmt.Println("r.Len() =", r.Len(), n, "bytes read.")
	fmt.Println("buf.Len() =", len(buf))

	fmt.Println("-------ReadRune&Reset-----------")
	s = "飞01哥2345"
	r = strings.NewReader(s)
	rr, _, _ := r.ReadRune()
	fmt.Println("rr", string(rr))
	rr, _, _ = r.ReadRune()
	fmt.Println("rr", string(rr))
	s1 := "123456789"
	r.Reset(s1)
	rr, _, _ = r.ReadRune()
	fmt.Println("rr", string(rr))

	fmt.Println("-------Seek&Size-----------")
	s = "0123456789"
	r = strings.NewReader(s)
	fmt.Println("r.Size() =", r.Size())
	r.Seek(1, io.SeekStart)
	c, _ := r.ReadByte()
	fmt.Println(string(c))
	fmt.Println("r.Size() =", r.Size())

	fmt.Println("-------UnreadByte&UnreadRune-----------")
	s = "飞01哥2345"
	r = strings.NewReader(s)
	rr, _, _ = r.ReadRune()
	fmt.Println("rr", string(rr))
	r.UnreadRune()
	rr, _, _ = r.ReadRune()
	fmt.Println("rr", string(rr))
	cc, _ := r.ReadByte()
	fmt.Println("cc", string(cc))
	cc, _ = r.ReadByte()
	r.UnreadByte()
	cc, _ = r.ReadByte()
	fmt.Println("cc", string(cc))

}
