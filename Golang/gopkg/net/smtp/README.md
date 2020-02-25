# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package smtp
## import "net/smtp"

Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321. It also implements the following extensions:

8BITMIME  RFC 1652
AUTH      RFC 2554
STARTTLS  RFC 3207
Additional extensions may be handled by clients.

The smtp package is frozen and is not accepting new features. Some external packages provide more functionality. See:

https://godoc.org/?q=smtp
Example
## Index
```
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
type Auth
func CRAMMD5Auth(username, secret string) Auth
func PlainAuth(identity, username, password, host string) Auth
type Client
func Dial(addr string) (*Client, error)
func NewClient(conn net.Conn, host string) (*Client, error)
func (c *Client) Auth(a Auth) error
func (c *Client) Close() error
func (c *Client) Data() (io.WriteCloser, error)
func (c *Client) Extension(ext string) (bool, string)
func (c *Client) Hello(localName string) error
func (c *Client) Mail(from string) error
func (c *Client) Noop() error
func (c *Client) Quit() error
func (c *Client) Rcpt(to string) error
func (c *Client) Reset() error
func (c *Client) StartTLS(config *tls.Config) error
func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)
func (c *Client) Verify(addr string) error
type ServerInfo
```
## Package Files

auth.go smtp.go

  