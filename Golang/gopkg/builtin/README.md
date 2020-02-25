# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package builtin
## import "builtin"
builtin package does not need to be imported.

package builtin provides documentation for go's predeclared identifiers. the items documented here are not actually in package builtin but their descriptions here allow godoc to present documentation for the language's special identifiers.

## index
```
Constants
func append(slice []Type, elems ...Type) []Type
func cap(v Type) int
func close(c chan<- Type)
func complex(r, i FloatType) ComplexType
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
func imag(c ComplexType) FloatType
func len(v Type) int
func make(t Type, size ...IntegerType) Type
func new(Type) *Type
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func real(c ComplexType) FloatType
func recover() interface{}
type ComplexType
type FloatType
type IntegerType
type Type
type Type1
type bool
type byte
type complex128
type complex64
type error
type float32
type float64
type int
type int16
type int32
type int64
type int8
type rune
type string
type uint
type uint16
type uint32
type uint64
type uint8
type uintptr
```

## Package Files
- builtin.go
  