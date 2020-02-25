# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package ioutil
## import "io/ioutil"  
Package ioutil implements some I/O utility functions.

## Index
```
Variables
func NopCloser(r io.Reader) io.ReadCloser  
func ReadAll(r io.Reader) ([]byte, error)  
func ReadDir(dirname string) ([]os.FileInfo, error)  
func ReadFile(filename string) ([]byte, error)  
func TempDir(dir, prefix string) (name string, err error)  
func TempFile(dir, pattern string) (f *os.File, err error)  
func WriteFile(filename string, data []byte, perm os.FileMode) error  
```
## Package Files
- ioutil.go
- tempfile.go
  