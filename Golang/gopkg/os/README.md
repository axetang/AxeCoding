# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package os
## import "os"

Package os provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.

The os interface is intended to be uniform across all operating systems. Features not generally available appear in the system-specific package syscall.

Here is a simple example, opening a file and reading some of it.

file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
If the open fails, the error string will be self-explanatory, like

open file.go: no such file or directory
The file's data can then be read into a slice of bytes. Read and Write take their byte counts from the length of the argument slice.

data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
## Index
```
Constants
Variables
func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
func Chtimes(name string, atime time.Time, mtime time.Time) error
func Clearenv()
func Environ() []string
func Executable() (string, error)
func Exit(code int)
func Expand(s string, mapping func(string) string) string
func ExpandEnv(s string) string
func Getegid() int
func Getenv(key string) string
func Geteuid() int
func Getgid() int
func Getgroups() ([]int, error)
func Getpagesize() int
func Getpid() int
func Getppid() int
func Getuid() int
func Getwd() (dir string, err error)
func Hostname() (name string, err error)
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPathSeparator(c uint8) bool
func IsPermission(err error) bool
func IsTimeout(err error) bool
func Lchown(name string, uid, gid int) error
func Link(oldname, newname string) error
func LookupEnv(key string) (string, bool)
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
func NewSyscallError(syscall string, err error) error
func Pipe() (r *File, w *File, err error)
func Readlink(name string) (string, error)
func Remove(name string) error
func RemoveAll(path string) error
func Rename(oldpath, newpath string) error
func SameFile(fi1, fi2 FileInfo) bool
func Setenv(key, value string) error
func Symlink(oldname, newname string) error
func TempDir() string
func Truncate(name string, size int64) error
func Unsetenv(key string) error
func UserCacheDir() (string, error)
type File
func Create(name string) (*File, error)
func NewFile(fd uintptr, name string) *File
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (f *File) Close() error
func (f *File) Fd() uintptr
func (f *File) Name() string
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) SetDeadline(t time.Time) error
func (f *File) SetReadDeadline(t time.Time) error
func (f *File) SetWriteDeadline(t time.Time) error
func (f *File) Stat() (FileInfo, error)
func (f *File) Sync() error
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)
type FileInfo
func Lstat(name string) (FileInfo, error)
func Stat(name string) (FileInfo, error)
type FileMode
func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool
func (m FileMode) Perm() FileMode
func (m FileMode) String() string
type LinkError
func (e *LinkError) Error() string
type PathError
func (e *PathError) Error() string
func (e *PathError) Timeout() bool
type ProcAttr
type Process
func FindProcess(pid int) (*Process, error)
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
func (p *Process) Kill() error
func (p *Process) Release() error
func (p *Process) Signal(sig Signal) error
func (p *Process) Wait() (*ProcessState, error)
type ProcessState
func (p *ProcessState) Exited() bool
func (p *ProcessState) Pid() int
func (p *ProcessState) String() string
func (p *ProcessState) Success() bool
func (p *ProcessState) Sys() interface{}
func (p *ProcessState) SysUsage() interface{}
func (p *ProcessState) SystemTime() time.Duration
func (p *ProcessState) UserTime() time.Duration
type Signal
type SyscallError
func (e *SyscallError) Error() string
func (e *SyscallError) Timeout() bool
```

## Package Files
- dir.go 
- dir_unix.go 
- env.go 
- error.go 
- error_posix.go 
- error_unix.go 
- exec.go 
- exec_posix.go 
- exec_unix.go 
- executable.go 
- executable_procfs.go 
- file.go file_posix.go 
- file_unix.go getwd.go 
- path.go path_unix.go 
- pipe_linux.go proc.go 
- stat.go stat_linux.go 
- stat_unix.go 
- sticky_notbsd.go 
- str.go sys.go 
- sys_linux.go 
- sys_unix.go 
- types.go 
- types_unix.go 
- wait_waitid.go

