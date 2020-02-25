# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package zip
## import "archive/zip"

Package zip provides support for reading and writing ZIP archives.

See: https://www.pkware.com/appnote

This package does not support disk spanning.

A note about ZIP64:

To be backwards compatible the FileHeader has both 32 and 64 bit Size fields. The 64 bit fields will always contain the correct value and for normal archives both fields will be the same. For files requiring the ZIP64 format the 32 bit fields will be 0xffffffff and the 64 bit fields must be used instead.

## Index
```
Constants
Variables
func RegisterCompressor(method uint16, comp Compressor)
func RegisterDecompressor(method uint16, dcomp Decompressor)
type Compressor
type Decompressor
type File
func (f *File) DataOffset() (offset int64, err error)
func (f *File) Open() (io.ReadCloser, error)
type FileHeader
func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
func (h *FileHeader) FileInfo() os.FileInfo
func (h *FileHeader) ModTime() time.Time
func (h *FileHeader) Mode() (mode os.FileMode)
func (h *FileHeader) SetModTime(t time.Time)
func (h *FileHeader) SetMode(mode os.FileMode)
type ReadCloser
func OpenReader(name string) (*ReadCloser, error)
func (rc *ReadCloser) Close() error
type Reader
func NewReader(r io.ReaderAt, size int64) (*Reader, error)
func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
type Writer
func NewWriter(w io.Writer) *Writer
func (w *Writer) Close() error
func (w *Writer) Create(name string) (io.Writer, error)
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
func (w *Writer) Flush() error
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
func (w *Writer) SetComment(comment string) error
func (w *Writer) SetOffset(n int64)
```
## Package Files

reader.go register.go struct.go writer.go

