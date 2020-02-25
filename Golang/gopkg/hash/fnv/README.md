# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package fnv
## import "hash/fnv"

Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. See https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.

All the hash.Hash implementations returned by this package also implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

## Index
```
func New128() hash.Hash
func New128a() hash.Hash
func New32() hash.Hash32
func New32a() hash.Hash32
func New64() hash.Hash64
func New64a() hash.Hash64
```
## Package Files

fnv.go