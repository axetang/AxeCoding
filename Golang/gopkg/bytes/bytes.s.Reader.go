/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Reader
**Type: struct
------------------------------------------------------------------------------------
**Definition:
// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a byte slice.
// Unlike a Buffer, a Reader is read-only and supports seeking.
// The zero value for Reader operates like a Reader of an empty slice.
type Reader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
func NewReader(b []byte) *Reader
func (r *Reader) Len() int
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error)
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) Reset(b []byte)
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) Size() int64
func (r *Reader) UnreadByte() error
func (r *Reader) UnreadRune() error
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
------------------------------------------------------------------------------------
**Description:
A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner,
and io.RuneScanner interfaces by reading from a byte slice. Unlike a Buffer, a Reader
is read-only and supports seeking.
NewReader returns a new Reader reading from b.
Len returns the number of bytes of the unread portion of the slice.
Read implements the io.Reader interface.
ReadAt implements the io.ReaderAt interface.
ReadByte implements the io.ByteReader interface.
ReadRune implements the io.RuneReader interface.
Reset resets the Reader to be reading from b.
Seek implements the io.Seeker interface.
Size returns the original length of the underlying byte slice. Size is the number
of bytes available for reading via ReadAt. The returned value is always the same and
is not affected by calls to any other method.
UnreadByte complements ReadByte in implementing the io.ByteScanner interface.
UnreadRune complements ReadRune in implementing the io.RuneScanner interface.
WriteTo implements the io.WriterTo interface.
------------------------------------------------------------------------------------
**要点总结:
0. Reader结构体实现了io.Reader，io.ReaderAt，io.WriterTo，io.Seeker, io.ByteScanner和
io.RuneScanner接口；底层是字节slice。Reader和Buffer不同，Reader是只读的；
1. Len()方法返回Reader底层字节切片的未读字节数；
2. Size方法返回Reader底层字节切片的初始长度；
3. Reset方法重置Reader底层字节切片为参数b([]byte)，Reset会初始化b的内容，因此这个方法执行后，
r.Len()=r.Size()
4. Read方法读取参数p（[]byte）的len(p)字节的内容，如果可读内容不足len(p)，返回已经读取的内容但
并不返回错误；如果已经无可读取内容，则返回EOF；
5. ReadByte读取并返回一个字节；
6. ReadRune读取一个rune，返回该rune及其字节数；
7. ReadAt从Reader中偏移为off字节的位置读取len(b)个字节，并返回已读取的字节数；如果遇到错误则返回
错误（通常是io.EOF）。ReadAt读取的数据不会从Reader中清除,即不会影响当前读取位置;
8. Seek方法实现了io.Seeker接口，用于设置下次读或写操作的位置，返回新的偏移位置字节数和错误。
	- whence的含义为：
		- 0：从起始位置计算offset
		- 1: 从当前位置计算offset
		- 2：从尾部位置计算offset
9.UnreadByte取消上次读取的最后一个字节，如果Buffer没有被读取过则UnreadByte也返回一个错误;
10. UnreadRune取消上次ReadRune读取的Unicode字符；如果上次读取操作不是ReadRune，则返回错误;
11. WriteTo实现了io.WriterTo接口，查阅源码，实现如下。WriteTo从r（*Reader）中读取所有剩余字节，
并写入参数w（io.Writer）。
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
	r.prevRune = -1
	if r.i >= int64(len(r.s)) {
		return 0, nil
	}
	b := r.s[r.i:]
	m, err := w.Write(b)
	if m > len(b) {
		panic("bytes.Reader.WriteTo: invalid Write count")
	}
	r.i += int64(m)
	n = int64(m)
	if m != len(b) && err == nil {
		err = io.ErrShortWrite
	}
	return
}
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"os"
)

//TestLenSizeReset  ddd
func TestLenSizeReset() {
	r := bytes.NewReader([]byte("12345"))
	fmt.Println("r.Len()", r.Len(), "r.Size()", r.Size())

	var buff [3]byte
	r.Read(buff[:])
	fmt.Println("r.Len()", r.Len(), "r.Size()", r.Size())
	var bufff [9]byte
	r.Reset(bufff[:])
	fmt.Println("r.Len()", r.Len(), "r.Size()", r.Size())

}

//TestRead ddd
func TestRead() {
	b := bytes.NewReader([]byte("12345"))
	var buff [3]byte
	n, err := b.Read(buff[:])
	fmt.Println(n, err, string(buff[:n]))
	n, err = b.Read(buff[:])
	fmt.Println(n, err, string(buff[:n]))
	n, err = b.Read(buff[:])
	fmt.Println(n, err, string(buff[:n]))
}

//TestReadByteReadRune ddd
func TestReadByteReadRune() {
	b := bytes.NewReader([]byte("12345"))
	c, err := b.ReadByte()
	fmt.Println(string(c), err)

	bb := bytes.NewReader([]byte("你好世界"))
	r, size, err := bb.ReadRune()
	fmt.Println(string(r), size, err)
}

//TestReadAtSeek ddd
func TestReadAtSeek() {
	b := bytes.NewReader([]byte("12345"))
	var buff [4]byte
	n, err := b.ReadAt(buff[:], 2)
	fmt.Println(n, err, string(buff[:n]))
	n, err = b.ReadAt(buff[:], 1)
	fmt.Println(n, err, string(buff[:n]))
	n, err = b.Read(buff[:])
	fmt.Println(n, err, string(buff[:n]))
	n, err = b.ReadAt(buff[:], 1)
	fmt.Println(n, err, string(buff[:n]))

	println()
	var data = []byte("123456")
	fmt.Println("seekhead:")
	bb := bytes.NewReader(data)
	// 把位置移到起始位置+1字节处，即0+1==1，对应数据为'2'
	fmt.Println(bb.Seek(1, 0))
	c, _ := bb.ReadByte()
	fmt.Println(string(c))
}

//TestUnreadByteUnreadRune ddd
func TestUnreadByteUnreadRune() {
	b := bytes.NewReader([]byte("12345"))
	fmt.Println(b.UnreadByte())
	var buff [6]byte
	n, _ := b.Read(buff[:])
	fmt.Printf("read: %s\n", string(buff[:n]))
	fmt.Println(b.UnreadByte())
	n, _ = b.Read(buff[:])
	fmt.Printf("read: %s\n", string(buff[:n]))

	println()

	bb := bytes.NewReader([]byte("你好世界"))
	var bufff [3]byte
	// 用Read读取第一个Unicode字符'你'
	bb.Read(bufff[:])
	fmt.Println(bb.UnreadRune())
	// 用ReadRuen读取第二个Unicode字符'好'
	r, _, _ := bb.ReadRune()
	fmt.Println(string(r))
	// 取消读取字符'好'
	fmt.Println(bb.UnreadRune())
	// 重新读取字符'好'
	r, _, _ = bb.ReadRune()
	fmt.Println(string(r))
}

// TestWriteTo dd
func TestWriteTo() {
	b := bytes.NewReader([]byte("1234567890\n"))
	b.WriteTo(os.Stdout)

}

func main() {
	/*fmt.Println("------------Len&Size&Reset------------")
	TestLenSizeReset()
	fmt.Println("------------Read------------")
	TestRead()
	fmt.Println("------------ReadByte&ReadRune------------")
	TestReadByteReadRune()
	fmt.Println("------------ReadAt&Seek------------")
	TestReadAtSeek()
	fmt.Println("------------UnreadByte&UnreadRune------------")
	TestUnreadByteUnreadRune()*/
	fmt.Println("------------WriteTo------------")
	TestWriteTo()
}
