# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package strconv
## import "strconv"
实现字符串和基本数据类型之间转换
Package strconv implements conversions to and from string representations of basic data types.
Numeric Conversions
The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
```
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
```
These assume decimal and the Go int type.

ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
```
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
```
The parse functions return the widest type (float64, int64, and uint64), but if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:

```
s := "2147483647" // biggest int32
i64, err := strconv.ParseInt(s, 10, 32)
...
i := int32(i64)
````
FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
```
s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
```
AppendBool, AppendFloat, AppendInt, and AppendUint are similar but append the formatted value to a destination slice.

String Conversions

Quote and QuoteToASCII convert strings to quoted Go string literals. The latter guarantees that the result is an ASCII string, by escaping any non-ASCII Unicode with \u:
```
q := Quote("Hello, 世界")
q := QuoteToASCII("Hello, 世界")
```
QuoteRune and QuoteRuneToASCII are similar but accept runes and return quoted Go rune literals.

Unquote and UnquoteChar unquote Go string and rune literals.

##Index
```
Constants
Variables
func AppendBool(dst []byte, b bool) []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
func AppendInt(dst []byte, i int64, base int) []byte
func AppendQuote(dst []byte, s string) []byte
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteToGraphic(dst []byte, s string) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
func Atoi(s string) (int, error)
func CanBackquote(s string) bool
func FormatBool(b bool) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func IsGraphic(r rune) bool
func IsPrint(r rune) bool
func Itoa(i int) string
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
func Quote(s string) string
func QuoteRune(r rune) string
func QuoteRuneToASCII(r rune) string
func QuoteRuneToGraphic(r rune) string
func QuoteToASCII(s string) string
func QuoteToGraphic(s string) string
func Unquote(s string) (string, error)
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
type NumError
func (e *NumError) Error() string
Examples

AppendBool
AppendFloat
AppendInt
AppendQuote
AppendQuoteRune
AppendQuoteRuneToASCII
AppendQuoteToASCII
AppendUint
Atoi
CanBackquote
FormatBool
FormatFloat
FormatInt
FormatUint
IsPrint
Itoa
NumError
ParseBool
ParseFloat
ParseInt
ParseUint
Quote
QuoteRune
QuoteRuneToASCII
QuoteToASCII
Unquote
UnquoteChar
```
## Package Files ¶
- atob.go 
- atof.go 
- atoi.go 
- decimal.go 
- doc.go 
- extfloat.go 
- ftoa.go 
- isprint.go 
- itoa.go 
- quote.go