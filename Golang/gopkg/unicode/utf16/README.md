# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package utf16
## import "unicode/utf16"

Package utf16 implements encoding and decoding of UTF-16 sequences.

## Index
```
func Decode(s []uint16) []rune
func DecodeRune(r1, r2 rune) rune
func Encode(s []rune) []uint16
func EncodeRune(r rune) (r1, r2 rune)
func IsSurrogate(r rune) bool
```

## Package Files
- utf16.go
  