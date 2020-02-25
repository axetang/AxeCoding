/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.ComplexType,builtin.FloatType,builtin.IntegerType,builtin.Type,
 builtin.Type1,builtin.bool,builtin.byte,builtin.complex128,builtin.complex64,
 builtin.float32,builtin.float64,builtin.int,builtin.int16,builtin.int32,builtin.int64,
 builtin.int8,builtin.rune,builtin.string,builtin.uint,builtin.uint16,builtin.uint32,
 builtin.uint64,builtin.uint8,builtin.uintptr
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type ComplexType complex64
 type FloatType float32
 type IntegerType int
 type Type int
 type Type1 int
 type bool bool
 type byte = uint8
 type complex128 complex128
 type complex64 complex64
 type error interface {
    Error() string
 }
 type float32 float32
 type float64 float64
 type int int
 type int16 int16
 type int32 int32
 type int64 int64
 type int8 int8
 type rune = int32
 type string string
 type uint uint
 type uint16 uint16
 type uint32 uint32
 type uint64 uint64
 type uint8 uint8
 type uintptr uintptr
 ------------------------------------------------------------------------------------
 **Description:
 ComplexType is here for the purposes of documentation only. It is a stand-in for
 either complex type: complex64 or complex128.
 FloatType is here for the purposes of documentation only. It is a stand-in for either
 float type: float32 or float64.
 IntegerType is here for the purposes of documentation only. It is a stand-in for any
 integer type: int, uint, int8 etc.
 Type is here for the purposes of documentation only. It is a stand-in for any Go
 type, but represents the same type for any given function invocation.
 nil is a predeclared identifier representing the zero value for a pointer, channel,
 func, interface, map, or slice type.
 Type1 is here for the purposes of documentation only. It is a stand-in for any Go
 type, but represents the same type for any given function invocation.
 bool is the set of boolean values, true and false.
 byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by
 convention, to distinguish byte values from 8-bit unsigned integer values.
 complex128 is the set of all complex numbers with float64 real and imaginary parts.
 complex64 is the set of all complex numbers with float32 real and imaginary parts.
 float32 is the set of all IEEE-754 32-bit floating-point numbers.
 float64 is the set of all IEEE-754 64-bit floating-point numbers.
 int is a signed integer type that is at least 32 bits in size. It is a distinct type,
 however, and not an alias for, say, int32.
 int16 is the set of all signed 16-bit integers. Range: -32768 through 32767.
 int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.
 int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through
 9223372036854775807.
 int8 is the set of all signed 8-bit integers. Range: -128 through 127.
 rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by
 convention, to distinguish character values from integer values.
 string is the set of all strings of 8-bit bytes, conventionally but not necessarily
 representing UTF-8-encoded text. A string may be empty, but not nil. Values of string
 type are immutable.
 uint is an unsigned integer type that is at least 32 bits in size. It is a distinct
 type, however, and not an alias for, say, uint32.
 uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.
 uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.
 uint64 is the set of all unsigned 64-bit integers. Range: 0 through
 18446744073709551615.
 uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.
 uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ComplexType在此只用作文档目的。它代表所有的复数类型：即complex64或complex128;
 2. FloatType在此只用作文档目的。它代表所有的浮点数类型：即float64或float32;
 3. IntType在此只用作文档目的。它代表所有的整数类型：即float64或float32;
 4. Type 在此只用作文档目的。 它代表所有Go的类型，但对于任何给定的函数请求来说，它都代表与其相同的类型;
 5. Type1在此只用作文档目的。 它代表所有Go的类型，但对于任何给定的函数请求来说，它都代表与其相同的类型;
 6. bool是布尔值的集合，即 true 和 false；
 7. byte为uint8的别名，它完全等价于uint8，习惯上用它来区别字节值和8位无符号整数值；
 8. complex64 是所有实部和虚部为float32的复数集合；
 9. complex128 是所有实部和虚部为float64的复数集合；
10. float32是所有IEEE-754 32位浮点数的集合；
11. float64是所有IEEE-754 64位浮点数的集合；
12. int 是带符号整数类型，其大小至少为32位。 它是一种确切的类型，而不是 int32 的别名；
13. int8是所有带符号8位整数的集合。范围：-128至 127；
14. int16是所有带符号16位整数的集合。范围：-32768 至32767；
15. int32是所有带符号32位整数的集合。范围: -2147483648至2147483647；
16. int64是所有带符号64位整数的集合。范围：-9223372036854775808至9223372036854775807；
17. rune为int32的别名，它完全等价于int32。习惯上用它来区别字符值和整数值；
18. string是所有8位字节的字符串集合，习惯上用于代表以UTF-8编码的文本，但并不必须如此。 string
可为空，但不为nil，string类型的值是不变的；
19. uint是无符号整数类型，其大小至少为32位。它是一种确切的类型，而不是uint32的别名;
20. uint8是所有无符号8位整数的集合。范围：0至255;
21. uint16是所有无符号16位整数的集合。范围：0至65535;
22. uint32是所有无符号32位整数的集合。范围：0至4294967295;
23. uint64是所有无符号64位整数的集合。范围：0至18446744073709551615;
24. uintptr为整数类型，其大小足以容纳任何指针的位模式。
*************************************************************************************/
package main

func main() {
}
