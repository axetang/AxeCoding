# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package tar
## import "archive/tar"

Package tar implements access to tar archives.

Tape archives (tar) are a file format for storing a sequence of files that can be read and written in a streaming manner. This package aims to cover most variations of the format, including those produced by GNU and BSD tar tools.

## Index
```
Constants
Variables
type Format
func (f Format) String() string
type Header
func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
func (h *Header) FileInfo() os.FileInfo
type Reader
func NewReader(r io.Reader) *Reader
func (tr *Reader) Next() (*Header, error)
func (tr *Reader) Read(b []byte) (int, error)
type Writer
func NewWriter(w io.Writer) *Writer
func (tw *Writer) Close() error
func (tw *Writer) Flush() error
func (tw *Writer) Write(b []byte) (int, error)
func (tw *Writer) WriteHeader(hdr *Header) error
```
## Package Files

common.go format.go reader.go stat_actime1.go stat_unix.go strconv.go writer.go

