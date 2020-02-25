/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Lchown
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Lchown(name string, uid, gid int) error
 ------------------------------------------------------------------------------------
 **Description:
 Lchown changes the numeric uid and gid of the named file. If the file is a symbolic
 link, it changes the uid and gid of the link itself. If there is an error, it will
 be of type *PathError.
 On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Lchown函数修改一个文件的所属用户和用户组;
 2. 如果文件为一个链接，修改的则是链接文件本身.
*************************************************************************************/
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fi, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)

	err = os.Lchown("./_iofiles/Hello.go", 502, -1) //nobody, nobody
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Println(os.IsPermission(err))
		return
	}

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)
}
