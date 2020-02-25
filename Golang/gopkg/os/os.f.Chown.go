/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Chown
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Chown(name string, uid, gid int) error
  ------------------------------------------------------------------------------------
 **Description:
 Chown changes the numeric uid and gid of the named file. If the file is a symbolic
 link, it changes the uid and gid of the link's target. A uid or gid of -1 means to
 not change that value. If there is an error, it will be of type *PathError.
     On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or EPLAN9 error,
     wrapped in *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Chown函数修改name指定的文件的uid和gid；
 2. 如果name指定的文件是符号链接，它修改链接指向的文件；
 3. 如果出错，返回*PathError；
 4. 根据执行来看，要修改文件的uid和gid，需要获得root权限或更改为root用户执行；
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

	err = os.Chown("./_iofiles/Hello.go", 502, -1) //nobody, nobody
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)
}
