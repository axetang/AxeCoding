# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package exec
## import "os/exec"

Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Unlike the "system" library call from C and other languages, the os/exec package intentionally does not invoke the system shell and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

Note that the examples in this package assume a Unix system. They may not run on Windows, and they do not run in the Go Playground used by golang.org and godoc.org.

## Index
```
Variables
func LookPath(file string) (string, error)
type Cmd
func Command(name string, arg ...string) *Cmd
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
func (c *Cmd) CombinedOutput() ([]byte, error)
func (c *Cmd) Output() ([]byte, error)
func (c *Cmd) Run() error
func (c *Cmd) Start() error
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
func (c *Cmd) Wait() error
type Error
func (e *Error) Error() string
type ExitError
func (e *ExitError) Error() string
```
## Package Files
- exec.go 
- exec_unix.go 
- lp_unix.go


  