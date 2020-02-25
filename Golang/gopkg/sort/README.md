# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package sort
## import "sort"

Package sort provides primitives for sorting slices and user-defined collections.

## Index
```
func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func Ints(a []int)
func IntsAreSorted(a []int) bool
func IsSorted(data Interface) bool
func Search(n int, f func(int) bool) int
func SearchFloat64s(a []float64, x float64) int
func SearchInts(a []int, x int) int
func SearchStrings(a []string, x string) int
func Slice(slice interface{}, less func(i, j int) bool)
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
func SliceStable(slice interface{}, less func(i, j int) bool)
func Sort(data Interface)
func Stable(data Interface)
func Strings(a []string)
func StringsAreSorted(a []string) bool
type Float64Slice
func (p Float64Slice) Len() int
func (p Float64Slice) Less(i, j int) bool
func (p Float64Slice) Search(x float64) int
func (p Float64Slice) Sort()
func (p Float64Slice) Swap(i, j int)
type IntSlice
func (p IntSlice) Len() int
func (p IntSlice) Less(i, j int) bool
func (p IntSlice) Search(x int) int
func (p IntSlice) Sort()
func (p IntSlice) Swap(i, j int)
type Interface
func Reverse(data Interface) Interface
type StringSlice
func (p StringSlice) Len() int
func (p StringSlice) Less(i, j int) bool
func (p StringSlice) Search(x string) int
func (p StringSlice) Sort()
func (p StringSlice) Swap(i, j int)
```
## Package Files
search.go slice.go sort.go zfuncversion.go

