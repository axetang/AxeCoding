# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package path
## import "path"

Package path implements utility routines for manipulating slash-separated paths.

The path package should only be used for paths separated by forward slashes, such as the paths in URLs. This package does not deal with Windows paths with drive letters or backslashes; to manipulate operating system paths, use the path/filepath package.

## Index
```
Variables
func Base(path string) string
func Clean(path string) string
func Dir(path string) string
func Ext(path string) string
func IsAbs(path string) bool
func Join(elem ...string) string
func Match(pattern, name string) (matched bool, err error)
func Split(path string) (dir, file string)
```
## Package Files
- match.go
- path.go