# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package user
## import "os/user"

Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of resolving user and group ids to names. One is written in pure Go and parses /etc/passwd and /etc/group. The other is cgo-based and relies on the standard C library (libc) routines such as getpwuid_r and getgrnam_r.

When cgo is available, cgo-based (libc-backed) code is used by default. This can be overriden by using osusergo build tag, which enforces the pure Go implementation.

## Index
```
type Group
func LookupGroup(name string) (*Group, error)
func LookupGroupId(gid string) (*Group, error)
type UnknownGroupError
func (e UnknownGroupError) Error() string
type UnknownGroupIdError
func (e UnknownGroupIdError) Error() string
type UnknownUserError
func (e UnknownUserError) Error() string
type UnknownUserIdError
func (e UnknownUserIdError) Error() string
type User
func Current() (*User, error)
func Lookup(username string) (*User, error)
func LookupId(uid string) (*User, error)
func (u *User) GroupIds() ([]string, error)
```
## Package Files
- cgo_lookup_unix.go 
- getgrouplist_unix.go 
- listgroups_unix.go 
- lookup.go 
- user.go
  