# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package list
## import "container/list"
Package list implements a doubly linked list.
To iterate over a list (where l is a *List):
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}
## Index
```
type Element
func (e *Element) Next() *Element
func (e *Element) Prev() *Element
type List
func New() *List
func (l *List) Back() *Element
func (l *List) Front() *Element
func (l *List) Init() *List
func (l *List) InsertAfter(v interface{}, mark *Element) *Element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) Len() int
func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)
func (l *List) PushBack(v interface{}) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v interface{}) *Element
func (l *List) PushFrontList(other *List)
func (l *List) Remove(e *Element) interface{}
```

## Package Files
list.go

