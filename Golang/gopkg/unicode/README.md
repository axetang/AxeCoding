# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package unicode
## import "unicode"
Package unicode provides data and functions to test some properties of Unicode code points.

## Index  
```
Constants
Variables
func In(r rune, ranges ...*RangeTable) bool
func Is(rangeTab *RangeTable, r rune) bool
func IsControl(r rune) bool
func IsDigit(r rune) bool
func IsGraphic(r rune) bool
func IsLetter(r rune) bool
func IsLower(r rune) bool
func IsMark(r rune) bool
func IsNumber(r rune) bool
func IsOneOf(ranges []*RangeTable, r rune) bool
func IsPrint(r rune) bool
func IsPunct(r rune) bool
func IsSpace(r rune) bool
func IsSymbol(r rune) bool
func IsTitle(r rune) bool
func IsUpper(r rune) bool
func SimpleFold(r rune) rune
func To(_case int, r rune) rune
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune
type CaseRange
type Range16
type Range32
type RangeTable
type SpecialCase
func (special SpecialCase) ToLower(r rune) rune
func (special SpecialCase) ToTitle(r rune) rune
func (special SpecialCase) ToUpper(r rune) rune
Bugs
Examples

SimpleFold
SpecialCase
To
ToLower
ToTitle
ToUpper
package (Is)
```
## Package Files
- casetables.go 
- digit.go 
- graphic.go 
- letter.go 
- tables.go

