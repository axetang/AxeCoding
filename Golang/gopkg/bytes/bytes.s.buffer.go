/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bytes
 **Element: bytes.Buffer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // A Buffer is a variable-sized buffer of bytes with Read and Write methods.
 // The zero value for Buffer is an empty buffer ready to use.
 type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.

	// FIXME: it would be advisable to align Buffer to cachelines to avoid false
	// sharing.
 }
 func NewBuffer(buf []byte) *Buffer
 func NewBufferString(s string) *Buffer
 func (b *Buffer) Bytes() []byte
 func (b *Buffer) Cap() int
 func (b *Buffer) Grow(n int)
 func (b *Buffer) Len() int
 func (b *Buffer) Next(n int) []byte
 func (b *Buffer) Read(p []byte) (n int, err error)
 func (b *Buffer) ReadByte() (byte, error)
 func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
 func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
 func (b *Buffer) ReadRune() (r rune, size int, err error)
 func (b *Buffer) ReadString(delim byte) (line string, err error)
 func (b *Buffer) Reset()
 func (b *Buffer) String() string
 func (b *Buffer) Truncate(n int)
 func (b *Buffer) UnreadByte() error
 func (b *Buffer) UnreadRune() error
 func (b *Buffer) Write(p []byte) (n int, err error)
 func (b *Buffer) WriteByte(c byte) error
 func (b *Buffer) WriteRune(r rune) (n int, err error)
 func (b *Buffer) WriteString(s string) (n int, err error)
 func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 A Buffer is a variable-sized buffer of bytes with Read and Write methods. The
 zero value for Buffer is an empty buffer ready to use.
 NewBuffer creates and initializes a new Buffer using buf as its initial contents.
 The new Buffer takes ownership of buf, and the caller should not use buf after
 this call. NewBuffer is intended to prepare a Buffer to read existing data. It
 can also be used to size the internal buffer for writing. To do that, buf should
 have the desired capacity but a length of zero.In most cases, new(Buffer)
 (or just declaring a Buffer variable) is sufficient to initialize a Buffer.
 NewBufferString creates and initializes a new Buffer using string s as its initial
 contents. It is intended to prepare a buffer to read an existing string.
 Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
 The slice is valid for use only until the next buffer modification (that is, only
 until the next call to a method like Read, Write, Reset, or Truncate). The slice
 aliases the buffer content at least until the next buffer modification, so
 immediate changes to the slice will affect the result of future reads.
 Cap returns the capacity of the buffer's underlying byte slice, that is, the
 total space allocated for the buffer's data.
 Grow grows the buffer's capacity, if necessary, to guarantee space for another n
 bytes. After Grow(n), at least n bytes can be written to the buffer without
 another allocation. If n is negative, Grow will panic. If the buffer can't grow
 it will panic with ErrTooLarge.
 Len returns the number of bytes of the unread portion of the buffer; b.Len() ==
 len(b.Bytes()).
 Next returns a slice containing the next n bytes from the buffer, advancing the
 buffer as if the bytes had been returned by Read. If there are fewer than n bytes
 in the buffer, Next returns the entire buffer. The slice is only valid until the
 next call to a read or write method.
 Read reads the next len(p) bytes from the buffer or until the buffer is drained.
 The return value n is the number of bytes read. If the buffer has no data to
 return, err is io.EOF (unless len(p) is zero); otherwise it is nil.
 ReadByte reads and returns the next byte from the buffer. If no byte is
 available, it returns error io.EOF.
 ReadBytes reads until the first occurrence of delim in the input, returning a
 slice containing the data up to and including the delimiter. If ReadBytes
 encounters an error before finding a delimiter, it returns the data read before
 the error and the error itself (often io.EOF). ReadBytes returns err != nil if
 and only if the returned data does not end in delim.
 ReadFrom reads data from r until EOF and appends it to the buffer, growing the
 buffer as needed. The return value n is the number of bytes read. Any error
 except io.EOF encountered during the read is also returned. If the buffer
 becomes too large, ReadFrom will panic with ErrTooLarge.
 ReadRune reads and returns the next UTF-8-encoded Unicode code point from the
 buffer. If no bytes are available, the error returned is io.EOF. If the bytes
 are an erroneous UTF-8 encoding, it consumes one byte and returns U+FFFD, 1.
 ReadString reads until the first occurrence of delim in the input, returning a
 string containing the data up to and including the delimiter. If ReadString
 encounters an error before finding a delimiter, it returns the data read before
 the error and the error itself (often io.EOF). ReadString returns err != nil if
 and only if the returned data does not end in delim.
 Reset resets the buffer to be empty, but it retains the underlying storage for
 use by future writes. Reset is the same as Truncate(0).
 String returns the contents of the unread portion of the buffer as a string. If
 the Buffer is a nil pointer, it returns "<nil>".
 To build strings more efficiently, see the strings.Builder type.
 Truncate discards all but the first n unread bytes from the buffer but continues
 to use the same allocated storage. It panics if n is negative or greater than
 the length of the buffer.
 UnreadByte unreads the last byte returned by the most recent successful read
 operation that read at least one byte. If a write has happened since the last
 read, if the last read returned an error, or if the read read zero bytes,
 UnreadByte returns an error.
 UnreadRune unreads the last rune returned by ReadRune. If the most recent read
 or write operation on the buffer was not a successful ReadRune, UnreadRune
 returns an error. (In this regard it is stricter than UnreadByte, which will
 unread the last byte from any read operation.)
 Write appends the contents of p to the buffer, growing the buffer as needed.
 The return value n is the length of p; err is always nil. If the buffer becomes
 too large, Write will panic with ErrTooLarge.
 WriteByte appends the byte c to the buffer, growing the buffer as needed. The
 returned error is always nil, but is included to match bufio.Writer's WriteByte.
 If the buffer becomes too large, WriteByte will panic with ErrTooLarge.
 WriteRune appends the UTF-8 encoding of Unicode code point r to the buffer,
 returning its length and an error, which is always nil but is included to match
 bufio.Writer's WriteRune. The buffer is grown as needed; if it becomes too large,
 WriteRune will panic with ErrTooLarge.
 WriteString appends the contents of s to the buffer, growing the buffer as needed.
 The return value n is the length of s; err is always nil. If the buffer becomes
 too large, WriteString will panic with ErrTooLarge.
 WriteTo writes data to w until the buffer is drained or an error occurs. The
 return value n is the number of bytes written; it always fits into an int, but
 it is int64 to match the io.WriterTo interface. Any error encountered during the
 write is also returned.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. bytes包定义了一些操作byte slice的便利操作。因为字符串可以表示为[]byte，因此bytes包定义
 的函数、方法等和strings包很类似，所以讲解时会和strings包类似甚至可以直接参考；
 1. Bytes方法返回Buffer缓存中尚未读取的字节切片[]byte;
 2. Read方法从Buffer中读取参数p([]byte)的长度len(p)字节的内容。如果len(p)大于Buffer的缓存
 buf的未读字节长度，即遇到EOF,Read方法返回读取的字节数，并不返回错误；
 3. Cap方法返回Buffer结构体底层[]byte切片的容量cap值；
 4. Len方法返回Buffer结构体底层[]byte切片未读内容的字节数；
 5. Next方法返回从当前读取位置开始的后n个字节，如果n大于可读字节数，Next将所有剩余未读部分返回；
 6. ReadByte方法从Buffer结构体的底层buf读取一个字节，如果没有字节可读，返回EOF；
 7. ReadBytes方法从Buffer结构体的底层buf读取，首次遇到终止符（参数delim byte）时停止，并
 返回读取到的内容的字节切片；如果没有遇到delim终止符，则返回读取到的所有内容以及io.EOF错误；
 8. ReadRune读取并返回一个UTF-8字符及其size；
 9. ReadString读取Buffer结构体的底层缓存，遇到终止符delim时停止并返回包含终止符的字符串；
 10. Reset清空Buffer结构体的底层字符切片；
 11. Truncate清空第n（参数）个字节后的底层字符切片数据，Reset相当于Truncate(0);
 12. String方法返回结构体Buffer底层数组未读取部分，并以string返回。如果Buffer是nil，返回nil；
 如果未读部分为0字节，则返回空字符串；
 13. ReadFrom从r中读取数据写入Buffer直到r.Read返回io.EOF，除了io.EOF之外的错误会被
 ReadFrom返回。返回值n为已读取的字节数，err为r.Read返回的错误。如果读取的数据太大，
 ReadFrom会panic并产生ErrTooLarge异常;
 14. WriteTo把Buffer中的数据写入w直到缓存数据为空或者遇到错误，返回值n总是足够用int表示，使用
 int64类型是为了和io.WriterTo接口匹配。任何写入时遇到的错误都会被返回;
 15. Write向Buffer缓冲区中写入p（[]byte）的内容，根据需要扩展buf缓冲区；
 16. WriteByte向Buffer的缓冲区buf中写入一个字节；
 17. WriteRune向Buffer的缓冲区buf写入参数r rune；
 18. WriteString向Buffer的缓冲区buf写入字符串参数s string；
 19. UnreadByte将Buffer缓冲区的读位置回退一个字节；
 20. UnreadRune将Buffer缓冲区的读位置回退一个rune；
 *************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func TestBytesAndRead() {
	b := bytes.NewBuffer([]byte("0123456789"))
	unread := b.Bytes()
	fmt.Printf("unread type:%T\n", unread)
	fmt.Println(string(unread))
	for i, c := range unread {
		fmt.Printf("%s ", string(c))
		unread[i] = 'A' + c - '0'
	}
	println()
	var buff [15]byte
	fmt.Printf("buf len:%d,cap:%d\n", len(buff), cap(buff))
	n, err := b.Read(buff[:])
	fmt.Println(n, "bytes read", "error is ", err, string(buff[:]))
}

func TestCapAndLen() {
	b := bytes.NewBuffer([]byte("123456"))
	fmt.Println("b.Len():", b.Len(), "b.Cap()", b.Cap())
	var buff [3]byte
	b.Read(buff[:])
	fmt.Println("b.Len():", b.Len(), "b.Cap()", b.Cap())
}

func TestNextAndGrow() {
	b := bytes.NewBuffer([]byte("0123456789"))
	fmt.Println("b.Len():", b.Len(), "b.Cap()", b.Cap())
	var buff [3]byte
	b.Read(buff[:])
	fmt.Println(string(buff[:]))
	s := b.Next(6)
	fmt.Printf("s is %s\n", string(s))

	var buff1 [3]byte
	b.Read(buff1[:])
	fmt.Println(string(buff1[:]))

	b.Grow(17)
	fmt.Println("b.Len():", b.Len(), "b.Cap()", b.Cap())

}

//TestReadByteAndReadBytes ddd
func TestReadByteAndReadBytes() {

	b := bytes.NewBuffer([]byte("0123456789"))
	c, err := b.ReadByte()
	fmt.Println(string(c), err)
	fmt.Println(string(b.Bytes()))

	bb := bytes.NewBuffer([]byte("123,456,7890"))
	line, err := bb.ReadBytes(',')
	fmt.Println(string(line))
	fmt.Println(err)
	line, err = bb.ReadBytes(',')
	fmt.Println(string(line))
	fmt.Println(err)

	line, err = bb.ReadBytes('^')
	fmt.Println(string(line))
	fmt.Println(err)
}

//TestReadRuneAndReadString kkk
func TestReadRuneAndReadString() {
	b := bytes.NewBuffer([]byte("你好世界"))
	r, size, err := b.ReadRune()
	fmt.Println(string(r), size, err)

	b = bytes.NewBuffer([]byte("123*456"))
	line, err := b.ReadString('*')
	fmt.Println(line, err)
}

// TestResetAndTruncate hhh
func TestResetAndTruncate() {
	b := bytes.NewBuffer([]byte("123"))
	fmt.Println(b.Len())
	b.Reset()
	fmt.Println(b.Len())

	bb := bytes.NewBuffer([]byte("0123456789"))
	bb.Truncate(5)
	fmt.Println(bb.String())

}

//TestString ddd
func TestString() {
	var b *bytes.Buffer
	fmt.Println(b.String())
	b = bytes.NewBuffer([]byte("123"))
	fmt.Println(b.String())
	var buf [3]byte
	b.Read(buf[:])
	fmt.Println(b.String())

}

//TestReadFromAndWriteTo ddd
func TestReadFromAndWriteTo() {
	r := bytes.NewBuffer([]byte("123"))
	b := bytes.NewBuffer(nil)
	n, err := b.ReadFrom(r)
	fmt.Println(n, err)
	fmt.Println(string(b.Bytes()))

	bb := bytes.NewBuffer([]byte("12345"))
	w := bytes.NewBuffer(nil)
	n, err = bb.WriteTo(w)
	fmt.Println(n, err)
	fmt.Println(string(w.Bytes()))
}

//TestWriteAndWriteByte ddd
func TestWriteAndWriteByte() {

	b := bytes.NewBuffer(nil)
	b.WriteByte('1')
	b.WriteByte('2')
	b.WriteByte('3')
	fmt.Println(string(b.Bytes()))

	bb := bytes.NewBuffer(nil)
	fmt.Println(bb.Write([]byte("12345")))
	fmt.Println(string(bb.Bytes()))
}

//TestWriteRuneAndWriteString ddd
func TestWriteRuneAndWriteString() {
	b := bytes.NewBuffer(nil)
	b.WriteRune('你')
	b.WriteRune('好')
	fmt.Println(string(b.Bytes()))

	bb := bytes.NewBuffer(nil)
	bb.WriteString("你好世界")
	fmt.Println(string(bb.Bytes()))
}

//TestUnreadByteAndUnreadRune ddd
func TestUnreadByteAndUnreadRune() {

	b := bytes.NewBuffer([]byte("12345"))
	fmt.Println(string(b.Next(3)))
	b.UnreadByte()
	fmt.Println(b.String())

	bb := bytes.NewBuffer([]byte("你好世界"))
	r, _, _ := bb.ReadRune()
	fmt.Println(string(r))
	fmt.Println(string(bb.Bytes()))
	bb.UnreadRune()
	fmt.Println(string(bb.Bytes()))
}

func main() {
	fmt.Println("------------Bytes&Read------------")
	TestBytesAndRead()
	fmt.Println("------------Cap&Len------------")
	TestCapAndLen()
	fmt.Println("------------Next&Grow------------")
	TestNextAndGrow()
	fmt.Println("------------ReadByte&ReadBytes------------")
	TestReadByteAndReadBytes()
	fmt.Println("------------ReadRune&ReadString------------")
	TestReadRuneAndReadString()
	fmt.Println("------------Reset&Truncate------------")
	TestResetAndTruncate()
	fmt.Println("------------String------------")
	TestString()

	fmt.Println("------------ReadFrom&WriteTo------------")
	TestReadFromAndWriteTo()

	fmt.Println("------------Write&WriteByte------------")
	TestWriteAndWriteByte()

	fmt.Println("------------WriteRune&WriteString------------")
	TestWriteRuneAndWriteString()

	fmt.Println("------------UnreadByte&UnreadRune------------")
	TestUnreadByteAndUnreadRune()

}
