# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package log
## import "log"

Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error and prints the date and time of each logged message. Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.

## Index
```
Constants
func Fatal(v ...interface{})
func Fatalf(format string, v ...interface{})
func Fatalln(v ...interface{})
func Flags() int
func Output(calldepth int, s string) error
func Panic(v ...interface{})
func Panicf(format string, v ...interface{})
func Panicln(v ...interface{})
func Prefix() string
func Print(v ...interface{})
func Printf(format string, v ...interface{})
func Println(v ...interface{})
func SetFlags(flag int)
func SetOutput(w io.Writer)
func SetPrefix(prefix string)
type Logger
func New(out io.Writer, prefix string, flag int) *Logger
func (l *Logger) Fatal(v ...interface{})
func (l *Logger) Fatalf(format string, v ...interface{})
func (l *Logger) Fatalln(v ...interface{})
func (l *Logger) Flags() int
func (l *Logger) Output(calldepth int, s string) error
func (l *Logger) Panic(v ...interface{})
func (l *Logger) Panicf(format string, v ...interface{})
func (l *Logger) Panicln(v ...interface{})
func (l *Logger) Prefix() string
func (l *Logger) Print(v ...interface{})
func (l *Logger) Printf(format string, v ...interface{})
func (l *Logger) Println(v ...interface{})
func (l *Logger) SetFlags(flag int)
func (l *Logger) SetOutput(w io.Writer)
func (l *Logger) SetPrefix(prefix string)
func (l *Logger) Writer() io.Writer
```
## Package Files

log.go

