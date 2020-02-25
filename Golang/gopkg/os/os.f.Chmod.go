/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Chmod
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Chmod(name string, mode FileMode) error
  ------------------------------------------------------------------------------------
 **Description:
 Chmod changes the mode of the named file to mode. If the file is a symbolic link, it
 changes the mode of the link's target. If there is an error, it will be of type
 *PathError.
 A different subset of the mode bits are used, depending on the operating system.
     On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and ModeSticky are
     used.
     On Windows, the mode must be non-zero but otherwise only the 0200 bit (owner
     writable) of mode is used; it controls whether the file's read-only attribute is
     set or cleared. attribute. The other bits are currently unused. Use mode 0400 for
     a read-only file and 0600 for a readable+writable file.
     On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive, and
     ModeTemporary are used.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Chomd修改name指定的文件的mode到参数指定的mode FileMode；
 2. 如果文件是一个符号链接，它会修改符号链接所指向的目标文件的mode。w
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)

	err = os.Chmod("./_iofiles/Hello.go", 0777)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)
}
