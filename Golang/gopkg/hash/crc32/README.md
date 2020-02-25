# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package crc32
## import "hash/crc32"

Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

Polynomials are represented in LSB-first form also known as reversed representation.

See https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials for information.

# Index
```
Constants
Variables
func Checksum(data []byte, tab *Table) uint32
func ChecksumIEEE(data []byte) uint32
func New(tab *Table) hash.Hash32
func NewIEEE() hash.Hash32
func Update(crc uint32, tab *Table, p []byte) uint32
type Table
func MakeTable(poly uint32) *Table
```
## Package Files

crc32.go crc32_amd64.go crc32_generic.go