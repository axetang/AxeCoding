# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package filepath
## import "path/filepath"

Package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.

The filepath package uses either forward slashes or backslashes, depending on the operating system. To process paths such as URLs that always use forward slashes regardless of the operating system, see the path package.

## Index
```
Constants
Variables
func Abs(path string) (string, error)
func Base(path string) string
func Clean(path string) string
func Dir(path string) string
func EvalSymlinks(path string) (string, error)
func Ext(path string) string
func FromSlash(path string) string
func Glob(pattern string) (matches []string, err error)
func HasPrefix(p, prefix string) bool
func IsAbs(path string) bool
func Join(elem ...string) string
func Match(pattern, name string) (matched bool, err error)
func Rel(basepath, targpath string) (string, error)
func Split(path string) (dir, file string)
func SplitList(path string) []string
func ToSlash(path string) string
func VolumeName(path string) string
func Walk(root string, walkFn WalkFunc) error
type WalkFunc
```

## Package Files
- match.go 
- path.go 
- path_unix.go 
- symlink.go 
- symlink_unix.go