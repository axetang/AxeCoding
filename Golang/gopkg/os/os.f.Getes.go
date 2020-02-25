/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: os
**Element: os.Getegid,os.Getenv,os.Geteuid,os.Getgid,os.Getgroups,os.Getpagesize,
os.Getpid,os.Getppid,os.Getuid,os.Getwd
**Type: func
------------------------------------------------------------------------------------
**Definition:
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
------------------------------------------------------------------------------------
**Description:
Getegid returns the numeric effective group id of the caller.
On Windows, it returns -1.
Getenv retrieves the value of the environment variable named by the key. It returns
the value, which will be empty if the variable is not present. To distinguish between
an empty value and an unset value, use LookupEnv.
Geteuid returns the numeric effective user id of the caller.
On Windows, it returns -1.
Getgid returns the numeric group id of the caller.
On Windows, it returns -1.
Getgroups returns a list of the numeric ids of groups that the caller belongs to.
On Windows, it returns syscall.EWINDOWS. See the os/user package for a possible
alternative.
Getpagesize returns the underlying system's memory page size.
Getpid returns the process id of the caller.
Getppid returns the process id of the caller's parent.
Getuid returns the numeric user id of the caller.
On Windows, it returns -1.
Getwd returns a rooted path name corresponding to the current directory. If the
current directory can be reached via multiple paths (due to symbolic links), Getwd
may return any one of them.
------------------------------------------------------------------------------------
**要点总结:
1. Getegid函数返回调用者的有效用户组id，windows环境下返回-1;
2. Getenv函数返回参数key指定的环境变量值；
3. Geteuid函数返回调用者的有效用户id，windows环境下返回-1；
4. Getgid返回调用者的用户组id；，windows环境下返回-1；
5. Getgroups返回调用者所属组的id；
6. Getpagesize返回底层操作系统的pagesize；
7. Getpid返回调用者进程id；
8. Getppid返回返回调用者父进程id；
9. Getuid返回调用者的用户id，windows环境下返回-1；
10. Getwd返回当前工作目录绝对目录。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-------Getegid--------")
	fmt.Printf("%d\n", os.Getegid())

	fmt.Println("-------Getenv--------")
	fmt.Printf("%s\n", os.Getenv("LANG"))

	fmt.Println("-------Geteuid--------")
	fmt.Printf("%d\n", os.Geteuid())

	fmt.Println("-------Getgid--------")
	fmt.Printf("%d\n", os.Getgid())

	fmt.Println("-------Getgroups--------")
	ids, err := os.Getgroups()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%v\n", ids)

	fmt.Println("-------Getpagesize--------")
	fmt.Printf("%d\n", os.Getpagesize())

	fmt.Println("-------Getpid--------")
	fmt.Printf("%d\n", os.Getpid())

	fmt.Println("-------Getppid--------")
	fmt.Printf("%d\n", os.Getppid())

	fmt.Println("-------Getuid--------")
	fmt.Printf("%d\n", os.Getuid())

	fmt.Println("-------Getwd--------")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The current directory is: %s\n", pwd)

}
