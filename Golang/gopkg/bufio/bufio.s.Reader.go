/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.Reader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Reader struct {
    // contains filtered or unexported fields
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int
	lastRuneSize int
 }
 func NewReader(rd io.Reader) *Reader
 func NewReaderSize(rd io.Reader, size int) *Reader
 func (b *Reader) Buffered() int
 func (b *Reader) Discard(n int) (discarded int, err error)
 func (b *Reader) Read(p []byte) (n int, err error)
 func (b *Reader) ReadByte() (byte, error)
 func (b *Reader) ReadBytes(delim byte) ([]byte, error)
 func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
 func (b *Reader) ReadRune() (r rune, size int, err error)
 func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
 func (b *Reader) ReadString(delim byte) (string, error)
 func (b *Reader) Reset(r io.Reader)
 func (r *Reader) Size() int
 func (b *Reader) UnreadByte() error
 func (b *Reader) UnreadRune() error
 func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 0. Reader implements buffering for an io.Reader object.
 1. NewReader returns a new Reader whose buffer has the default size.
 2. NewReaderSize returns a new Reader whose buffer has at least the specified size.
 If the argument io.Reader is already a Reader with large enough size, it returns
 the underlying Reader.
 3. Buffered returns the number of bytes that can be read from the current buffer.
 4. Discard skips the next n bytes, returning the number of bytes discarded.
 If Discard skips fewer than n bytes, it also returns an error. If 0<=n<=b.Buffered(),
 Discard is guaranteed to succeed without reading from the underlying io.Reader.
 5. Peek returns the next n bytes without advancing the reader. The bytes stop
 being valid at the next read call. If Peek returns fewer than n bytes, it also
 returns an error explaining why the read is short. The error is ErrBufferFull if
 n is larger than b's buffer size.
 6. Read reads data into p. It returns the number of bytes read into p. The bytes
 are taken from at most one Read on the underlying Reader, hence n may be less
 than len(p). At EOF, the count will be zero and err will be io.EOF.
 7. ReadByte reads and returns a single byte. If no byte is available, returns
 an error.
 8. ReadBytes reads until the first occurrence of delim in the input, returning
 a slice containing the data up to and including the delimiter. If ReadBytes
 encounters an error before finding a delimiter, it returns the data read before
 the error and the error itself (often io.EOF). ReadBytes returns err != nil if
 and only if the returned data does not end in delim. For simple uses, a Scanner
 may be more convenient.
 9. ReadLine is a low-level line-reading primitive. Most callers should use
 ReadBytes('\n') or ReadString('\n') instead or use a Scanner. ReadLine tries to
 return a single line, not including the end-of-line bytes. If the line was too
 long for the buffer then isPrefix is set and the beginning of the line is r
 eturned. The rest of the line will be returned from future calls. isPrefix will
 be false when returning the last fragment of the line. The returned buffer is
 only valid until the next call to ReadLine. ReadLine either returns a non-nil
 line or it returns an error, never both. The text returned from ReadLine does not
 include the line end ("\r\n" or "\n"). No indication or error is given if the
 input ends without a final line end. Calling UnreadByte after ReadLine will
 always unread the last byte read (possibly a character belonging to the line
 end) even if that byte is not part of the line returned by ReadLine.
 10. ReadRune reads a single UTF-8 encoded Unicode character and returns the rune
 and its size in bytes. If the encoded rune is invalid, it consumes one byte and
 returns unicode.ReplacementChar (U+FFFD) with a size of 1.
 11. ReadSlice reads until the first occurrence of delim in the input, returning
 a slice pointing at the bytes in the buffer. The bytes stop being valid at the
 next read. If ReadSlice encounters an error before finding a delimiter, it
 returns all the data in the buffer and the error itself (often io.EOF).
 ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
 Because the data returned from ReadSlice will be overwritten by the next I/O
 operation, most clients should use ReadBytes or ReadString instead. ReadSlice
 returns err != nil if and only if line does not end in delim.
 12. ReadString reads until the first occurrence of delim in the input, returning
 a string containing the data up to and including the delimiter. If ReadString
 encounters an error before finding a delimiter, it returns the data read before
 the error and the error itself (often io.EOF). ReadString returns err != nil if
 and only if the returned data does not end in delim. For simple uses, a Scanner
 may be more convenient.
 13. Reset discards any buffered data, resets all state, and switches the buffered
 reader to read from r.
 14. Size returns the size of the underlying buffer in bytes.
 15. UnreadByte unreads the last byte. Only the most recently read byte can be unread.
 16. UnreadRune unreads the last rune. If the most recent read operation on the
 buffer was not a ReadRune, UnreadRune returns an error. (In this regard it is
 stricter than UnreadByte, which will unread the last byte from any read operation.)
 17. WriteTo implements io.WriterTo. This may make multiple calls to the Read
 method of the underlying Reader. If the underlying reader supports the WriteTo
 method, this calls the underlying WriteTo without buffering.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Reader实现了对一个io.Reader对象的缓冲读；bufio包提供了两个实例化bufio.Reader对象的函数:
 NewReader和NewReaderSize。其中NewReader函数是调用NewReaderSize函数实现的;
 1. NewReader返回一个新的Reader，这个Reader的大小是默认的大小defaultBufSize=4096；
 2. NewReaderSize返回了一个新的读取器，这个读取器的缓存大小至少大于制定的大小。 如果io.Reader
 参数已经是一个有足够大缓存的读取器，它就会返回这个Reader了；
 3. Buffered返回当前缓存的可读字节数；
 4. Discard方法忽略并跳过下n个字节的读取，返回跳过的字节数；
 5. Peek返回尚未读取的下n个字节，它不会移动读取位置；
 6. Read读取数据到p并返回读取到p的字节数；
 7. ReadByte读取和返回一个单字节，如果没有字节可以读取返回error；
 8. ReadBytes持续按字节读取直到遇到终止符，它返回的byte slice包含从起始读取位置到终止符的内容
 （包括终止符）；如果ReadBytes在遇到终止符之前就捕获到一个错误，它就会返回遇到错误之前已经读取的
 数据，和这个捕获到的错误（经常是 io.EOF）;
 9. ReadLine是底层用于读取一行数据的方法。可以使用ReadBytes('\n')或者ReadString('\n')来代替。
 ReadLine试图返回一行，不包括结尾的回车字符。如果一行太长了（超过缓冲区长度，如下程序测试，缓冲区
 的长度为4096字节），isPrefix会设置为true并且只返回前面的数据，剩余的数据会在下一次的调用中返回。
 当返回最后一行数据时，isPrefix会设置为false；
 10. ReadRune读取一个UTF-8编码的字符，并将其对应的Unicode编码和所占字节数返回。如果编码错误，
 ReadRune只读取一个字节并返回unicode.ReplacementChar(U+FFFD)和长度1;
 11. ReadSlice读取数据直到delim终止符出现，并返回读取数据的字节切片；
 12. ReadString读取直到遇到终止符，返回的string包含从当前到终止符的内容（包括终止符）;
 13. UnreadByte将上次读取的最后一个字节标志为未读;
 14. UnreadRune将最后一个rune设置为未读。如果最新的在buffer上的操作不是ReadRune，
 则UnreadRune返回一个error;
 15. Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据;
 16. Size方法返回Reader的buffer的size字节数，从下面程序测试看是4096，哪里设置需要查看源码确认；
 17. WriteTo实现了io.WriterTo接口的WriteTo方法,其具体实现中调用了WriteBuf方法，具体实现还需
后续深入标准库源代码了解。
*************************************************************************************/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func TestBuffered() {
	rb := bytes.NewBuffer([]byte("012345678"))
	r := bufio.NewReader(rb)
	fmt.Println("Buffered is r.w-r.r: ", r.Buffered())
	var buf [4]byte
	r.Read(buf[:])
	fmt.Println(r.Buffered())
	r.Read(buf[:])
	fmt.Println(r.Buffered())
}

func TestPeek() {
	rb := bytes.NewBuffer([]byte("012345678"))
	r := bufio.NewReader(rb)
	b1, _ := r.Peek(4)
	fmt.Println(string(b1))
	b2, err := r.Peek(20)
	fmt.Println(string(b2), err)
}
func TestDiscard() {
	rb := bytes.NewBuffer([]byte("0123456789"))
	r := bufio.NewReader(rb)
	n1, err := r.Discard(4)
	fmt.Println(n1, err)
	var buf [4]byte
	r.Read(buf[:])
	fmt.Println(string(buf[0:4]))
	n2, err := r.Discard(2)
	fmt.Println(n2, err)
	n3, err := r.Discard(1)
	fmt.Println(n3, err)
}
func TestRead() {
	rb := bytes.NewBuffer([]byte("12345678"))
	r := bufio.NewReader(rb)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Printf("%d, %v\n", n, err)
	fmt.Println(string(buf[:n]))
}
func TestReadByte() {
	rb := bytes.NewBuffer([]byte("012345678"))
	r := bufio.NewReader(rb)
	for i := 0; i < 10; i++ {
		b, err := r.ReadByte()
		fmt.Printf("%c %v,", b, err)
	}
	fmt.Println()
}
func TestReadBytes() {
	rb := bytes.NewBuffer([]byte("0123456$789"))
	r := bufio.NewReader(rb)
	b, err := r.ReadBytes('$')
	fmt.Printf("%s, %v\n", string(b), err)
	rb = bytes.NewBuffer([]byte("0123456$789"))
	r = bufio.NewReader(rb)
	b, err = r.ReadBytes('^')
	fmt.Printf("%s, %v\n", string(b), err)
}
func TestReadLine() {
	sl := make([]byte, 4099)
	for i := 0; i < 4099; i++ {
		sl[i] = 'a'
	}

	rb := bytes.NewBuffer([]byte("123\r\n456"))
	r := bufio.NewReader(rb)
	line, prefix, err := r.ReadLine()
	if err == nil {
		fmt.Printf("%v, %s, %v\n", line, string(line), prefix)
	}
	println()

	rb = bytes.NewBuffer(sl[:])
	r = bufio.NewReader(rb)
	line, prefix, err = r.ReadLine()
	if err == nil {
		fmt.Printf("%v, %s, %v\n", line, string(line), prefix)
		fmt.Println("bytes of line read is", len(string(line)), "prefix is", prefix)
	}
	line, prefix, err = r.ReadLine()
	if err == nil {
		fmt.Printf("%v, %s, %v\n", line, string(line), prefix)
		fmt.Println("bytes of left line read is", len(string(line)), "prefix is", prefix)
	}

	println()
}

func TestReadRune() {
	b := []byte("你好世界")
	rb := bytes.NewBuffer(b)
	r := bufio.NewReader(rb)
	c, size, err := r.ReadRune()
	if err == nil {
		fmt.Println(string(c))
		fmt.Printf("%x, %d\n", c, size)
		fmt.Printf("%x\n", b[:size])
	}
}

func TestReadSlice() {
	rb := bytes.NewBuffer([]byte("12345678^9"))
	r := bufio.NewReader(rb)
	line, err := r.ReadSlice('$')
	if err != nil {
		fmt.Println("Error is", err.Error())
	}
	fmt.Println(string(line))
	line, err = r.ReadSlice('$')
	fmt.Println("$", string(line))
	line, err = r.ReadSlice('^')
	fmt.Println("^", string(line))
}

func TestReadString() {
	rb := bytes.NewBuffer([]byte("1234$5678"))
	r := bufio.NewReader(rb)
	line, err := r.ReadString('^')
	if err == nil {
		fmt.Println(line)
	} else {
		fmt.Println("error is ", err)
		fmt.Println(line)
	}
}

func TestUnreadByte() {
	rb := bytes.NewBuffer([]byte("1234,56$"))
	r := bufio.NewReader(rb)
	line, _ := r.ReadSlice(',')
	fmt.Println(string(line))
	// unread ','
	fmt.Println(r.UnreadByte())
	line, _ = r.ReadSlice('$')
	fmt.Println(string(line))
}

func TestUnreadRune() {
	rb := bytes.NewBuffer([]byte("123456"))
	r := bufio.NewReader(rb)
	r.ReadByte()
	// error occurs
	fmt.Println(r.UnreadRune())
	c, _, _ := r.ReadRune()
	fmt.Printf("read %s\n", string(c))
	// no error happens
	fmt.Println(r.UnreadRune())
	c, _, _ = r.ReadRune()
	fmt.Printf("read %s\n", string(c))
}

func TestReset() {
	rb := bytes.NewBuffer([]byte("1234$567^8"))
	r := bufio.NewReader(rb)
	line, err := r.ReadString('$')

	fmt.Println("error is ", err)
	fmt.Println(line)

	r.Reset(r)
	line, err = r.ReadString('^')
	fmt.Println(line)

}
func TestSize() {
	rb := bytes.NewBuffer([]byte("1234$567^8"))
	r := bufio.NewReader(rb)
	line, err := r.ReadString('$')

	fmt.Println("error is ", err)
	fmt.Println("ReadString $:", line, "r buf size:", r.Size())

	line, err = r.ReadString('^')
	fmt.Println("ReadString ^:", line, "r buf size:", r.Size())

}

func main() {
	/*fmt.Println("------Buffered-------------")
	TestBuffered()
	fmt.Println("--------Peek------------")
	TestPeek()
	fmt.Println("-----Discard---------------")
	TestDiscard()
	fmt.Println("-----Read---------------")
	TestRead()
	fmt.Println("-----ReadByte---------------")
	TestReadByte()
	fmt.Println("-----ReadBytes---------------")
	TestReadBytes()
	fmt.Println("-----ReadLine---------------")
	TestReadLine()
	fmt.Println("-----ReadRune---------------")
	TestReadRune()
	fmt.Println("-----ReadSlice---------------")
	TestReadSlice()
	fmt.Println("-----ReadString---------------")
	TestReadString()
	fmt.Println("-----UnreadByte---------------")
	TestUnreadByte()
	fmt.Println("-----UnreadRune---------------")
	TestUnreadRune()*/
	fmt.Println("-----Reset---------------")
	//TestReset()
	fmt.Println("-----Size---------------")
	TestSize()
	fmt.Println("----WriteTo---------------")
	//TestWriteTo()*/
}
