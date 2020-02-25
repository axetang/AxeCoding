# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  

# package adler32
## import "hash/adler32"

Package adler32 implements the Adler-32 checksum.

It is defined in RFC 1950:

Adler-32 is composed of two sums accumulated per byte: s1 is
the sum of all bytes, s2 is the sum of all s1 values. Both sums
are done modulo 65521. s1 is initialized to 1, s2 to zero.  The
Adler-32 checksum is stored as s2*65536 + s1 in most-
significant-byte first (network) order.
## Index
Constants
func Checksum(data []byte) uint32
func New() hash.Hash32
## Package Files

adler32.go

