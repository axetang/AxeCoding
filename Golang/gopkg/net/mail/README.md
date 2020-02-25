# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package mail
## import "net/mail"

Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by RFC 5322 and extended by RFC 6532. Notable divergences:

* Obsolete address formats are not parsed, including addresses with
  embedded route information.
* The full range of spacing (the CFWS syntax element) is not supported,
  such as breaking addresses across lines.
* No unicode normalization is performed.
* The special characters ()[]:;@\, are allowed to appear unquoted in names.
## Index
```
Variables
func ParseDate(date string) (time.Time, error)
type Address
func ParseAddress(address string) (*Address, error)
func ParseAddressList(list string) ([]*Address, error)
func (a *Address) String() string
type AddressParser
func (p *AddressParser) Parse(address string) (*Address, error)
func (p *AddressParser) ParseList(list string) ([]*Address, error)
type Header
func (h Header) AddressList(key string) ([]*Address, error)
func (h Header) Date() (time.Time, error)
func (h Header) Get(key string) string
type Message
func ReadMessage(r io.Reader) (msg *Message, err error)
```

## Package Files

message.go