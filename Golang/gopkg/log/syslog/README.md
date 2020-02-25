# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package syslog
## import "log/syslog"

Package syslog provides a simple interface to the system log service. It can send messages to the syslog daemon using UNIX domain sockets, UDP or TCP.

Only one call to Dial is necessary. On write failures, the syslog client will attempt to reconnect to the server and write again.

The syslog package is frozen and is not accepting new features. Some external packages provide more functionality. See:

https://godoc.org/?q=syslog
## Index
```
func NewLogger(p Priority, logFlag int) (*log.Logger, error)
type Priority
type Writer
func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)
func New(priority Priority, tag string) (*Writer, error)
func (w *Writer) Alert(m string) error
func (w *Writer) Close() error
func (w *Writer) Crit(m string) error
func (w *Writer) Debug(m string) error
func (w *Writer) Emerg(m string) error
func (w *Writer) Err(m string) error
func (w *Writer) Info(m string) error
func (w *Writer) Notice(m string) error
func (w *Writer) Warning(m string) error
func (w *Writer) Write(b []byte) (int, error)
Bugs
Examples
```

## Package Files

doc.go syslog.go syslog_unix.go)