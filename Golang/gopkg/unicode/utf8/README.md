# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package utf8
## import "unicode/utf8"

Package utf8 implements functions and constants to support text encoded in UTF-8. It includes functions to translate between runes and UTF-8 byte sequences.

## Index
```
Constants
func DecodeLastRune(p []byte) (r rune, size int)
func DecodeLastRuneInString(s string) (r rune, size int)
func DecodeRune(p []byte) (r rune, size int)
func DecodeRuneInString(s string) (r rune, size int)
func EncodeRune(p []byte, r rune) int
func FullRune(p []byte) bool
func FullRuneInString(s string) bool
func RuneCount(p []byte) int
func RuneCountInString(s string) (n int)
func RuneLen(r rune) int
func RuneStart(b byte) bool
func Valid(p []byte) bool
func ValidRune(r rune) bool
func ValidString(s string) bool

```
## Package Files
- utf8.go