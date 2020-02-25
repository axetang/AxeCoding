# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package crc64
## import "hash/crc64"

Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

## Index
```
Constants
func Checksum(data []byte, tab *Table) uint64
func New(tab *Table) hash.Hash64
func Update(crc uint64, tab *Table, p []byte) uint64
type Table
func MakeTable(poly uint64) *Table
```
## Package Files

crc64.go