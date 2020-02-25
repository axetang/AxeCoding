# package io
**标准包文档链接如下：**  
- [English Version](https://godoc.org/io)

- [中文链接](http://docscn.studygolang.com/pkg/io/)


# package io
## import "io"

Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should not assume they are safe for parallel execution.

## Index
```
Constants  
Variables  
func Copy(dst Writer, src Reader) (written int64, err error)  
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)  
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)  
func Pipe() (*PipeReader, *PipeWriter)  
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)  
func ReadFull(r Reader, buf []byte) (n int, err error)  
func WriteString(w Writer, s string) (n int, err error)  
type ByteReader  
type ByteScanner  
type ByteWriter  
type Closer  
type LimitedReader  
func (l *LimitedReader) Read(p []byte) (n int, err error)  
type PipeReader  
func (r *PipeReader) Close() error  
func (r *PipeReader) CloseWithError(err error) error  
func (r *PipeReader) Read(data []byte) (n int, err error)  
type PipeWriter  
func (w *PipeWriter) Close() error  
func (w *PipeWriter) CloseWithError(err error) error  
func (w *PipeWriter) Write(data []byte) (n int, err error)  
type ReadCloser  
type ReadSeeker  
type ReadWriteCloser  
type ReadWriteSeeker  
type ReadWriter  
type Reader  
func LimitReader(r Reader, n int64) Reader  
func MultiReader(readers ...Reader) Reader  
func TeeReader(r Reader, w Writer) Reader  
type ReaderAt  
type ReaderFrom  
type RuneReader  
type RuneScanner  
type SectionReader  
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader  
func (s *SectionReader) Read(p []byte) (n int, err error)  
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)  
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)  
func (s *SectionReader) Size() int64  
type Seeker  
type WriteCloser  
type WriteSeeker  
type Writer  
func MultiWriter(writers ...Writer) Writer  
type WriterAt  
type WriterTo  
```
## Package Files
- io.go
- multi.go 
- pipe.go