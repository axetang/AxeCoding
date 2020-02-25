# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package ring
## import "container/ring"

Package ring implements operations on circular lists.

## Index
```
type Ring
func New(n int) *Ring
func (r *Ring) Do(f func(interface{}))
func (r *Ring) Len() int
func (r *Ring) Link(s *Ring) *Ring
func (r *Ring) Move(n int) *Ring
func (r *Ring) Next() *Ring
func (r *Ring) Prev() *Ring
func (r *Ring) Unlink(n int) *Ring
```
## Package Files

ring.go

  