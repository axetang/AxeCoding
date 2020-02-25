/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type FileInfo interface {
     Name() string       // base name of the file
     Size() int64        // length in bytes for regular files; system-dependent for others
     Mode() FileMode     // file mode bits
     ModTime() time.Time // modification time
     IsDir() bool        // abbreviation for Mode().IsDir()
     Sys() interface{}   // underlying data source (can return nil)
 }

 func Lstat(name string) (FileInfo, error)
 func Stat(name string) (FileInfo, error)
 ------------------------------------------------------------------------------------
 **Description:
 A FileInfo describes a file and is returned by Stat and Lstat.
 Lstat returns a FileInfo describing the named file. If the file is a symbolic link,
 the returned FileInfo describes the symbolic link. Lstat makes no attempt to follow
 the link. If there is an error, it will be of type *PathError.
 Stat returns a FileInfo describing the named file. If there is an error, it will be
 of type *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Lstat函数获取一个文件的信息，如果这个文件是一个链接，返回的是链接本身的信息；
 2. Stat函数获取一个文件的信息，如果这个文件是一个链接，返回的是目标文件的信息。
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
	fmt.Printf("file info: %+v\n", fi)

	fi, err = os.Lstat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("file info: %+v\n", fi)
}
