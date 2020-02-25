# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package url
## import "net/url"

Package url parses URLs and implements query escaping.

## Index
```
func PathEscape(s string) string
func PathUnescape(s string) (string, error)
func QueryEscape(s string) string
func QueryUnescape(s string) (string, error)
type Error
func (e *Error) Error() string
func (e *Error) Temporary() bool
func (e *Error) Timeout() bool
type EscapeError
func (e EscapeError) Error() string
type InvalidHostError
func (e InvalidHostError) Error() string
type URL
func Parse(rawurl string) (*URL, error)
func ParseRequestURI(rawurl string) (*URL, error)
func (u *URL) EscapedPath() string
func (u *URL) Hostname() string
func (u *URL) IsAbs() bool
func (u *URL) MarshalBinary() (text []byte, err error)
func (u *URL) Parse(ref string) (*URL, error)
func (u *URL) Port() string
func (u *URL) Query() Values
func (u *URL) RequestURI() string
func (u *URL) ResolveReference(ref *URL) *URL
func (u *URL) String() string
func (u *URL) UnmarshalBinary(text []byte) error
type Userinfo
func User(username string) *Userinfo
func UserPassword(username, password string) *Userinfo
func (u *Userinfo) Password() (string, bool)
func (u *Userinfo) String() string
func (u *Userinfo) Username() string
type Values
func ParseQuery(query string) (Values, error)
func (v Values) Add(key, value string)
func (v Values) Del(key string)
func (v Values) Encode() string
func (v Values) Get(key string) string
func (v Values) Set(key, value string)
```

## Package Files

url.go  