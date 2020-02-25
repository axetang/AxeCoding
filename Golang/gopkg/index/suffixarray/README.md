# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)

# package suffixarray
## import "index/suffixarray"

Package suffixarray implements substring search in logarithmic time using an in-memory suffix array.

Example use:

// create index for some data
index := suffixarray.New(data)

// lookup byte slice s
offsets1 := index.Lookup(s, -1) // the list of all indices where s occurs in data
offsets2 := index.Lookup(s, 3)  // the list of at most 3 indices where s occurs in data
## Index
```
type Index
func New(data []byte) *Index
func (x *Index) Bytes() []byte
func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
func (x *Index) Lookup(s []byte, n int) (result []int)
func (x *Index) Read(r io.Reader) error
func (x *Index) Write(w io.Writer) error
```
## Package Files

qsufsort.go suffixarray.go
