# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package bufio 

## import "bufio"
Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
## Index
```
Constants
Variables
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
type ReadWriter
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
type Reader
func NewReader(rd io.Reader) *Reader
func NewReaderSize(rd io.Reader, size int) *Reader
func (b *Reader) Buffered() int
func (b *Reader) Discard(n int) (discarded int, err error)
func (b *Reader) Peek(n int) ([]byte, error)
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
type Scanner
func NewScanner(r io.Reader) *Scanner
func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string
type SplitFunc
type Writer
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
```
## Package Files
- bufio.go 
- scan.go