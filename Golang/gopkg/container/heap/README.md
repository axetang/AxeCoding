# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package heap
## import "container/heap"

Package heap provides heap operations for any type that implements heap.Interface. A heap is a tree with the property that each node is the minimum-valued node in its subtree.

The minimum element in the tree is the root, at index 0.

A heap is a common way to implement a priority queue. To build a priority queue, implement the Heap interface with the (negative) priority as the ordering for the Less method, so Push adds items while Pop removes the highest-priority item from the queue. The Examples include such an implementation; the file example_pq_test.go has the complete source.

## Index
```
func Fix(h Interface, i int)
func Init(h Interface)
func Pop(h Interface) interface{}
func Push(h Interface, x interface{})
func Remove(h Interface, i int) interface{}
type Interface
```

## Package Files

heap.go

