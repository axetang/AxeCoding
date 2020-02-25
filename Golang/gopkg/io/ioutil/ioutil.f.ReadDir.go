/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ioutil
 **Element: ioutil.ReadDir
 **Type:func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadDir(dirname string) ([]os.FileInfo, error)
 ------------------------------------------------------------------------------------
 **Description:
 ReadDir reads the directory named by dirname and returns a list of directory
 entries sorted by filename.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ReadDir读取dirname，返回按照文件名排序的目录列表，类型是[]os.FileInfo切片；
 2. ReadDir不会递归读取下层目录内容。
 3. os.FileInfo定义如下，此函数实际上返回的是os.fileStat结构体，它实现了os.FileInfo接口。
 // A FileInfo describes a file and is returned by Stat and Lstat.
 type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
 }
 type fileStat struct {
	name    string
	size    int64
	mode    FileMode
	modTime time.Time
	sys     syscall.Stat_t
}
func (fs *fileStat) Size() int64        { return fs.size }
func (fs *fileStat) Mode() FileMode     { return fs.mode }
func (fs *fileStat) ModTime() time.Time { return fs.modTime }
func (fs *fileStat) Sys() interface{}   { return &fs.sys }
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("./_pkgsrc/ioutil")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files type is %T,\nvalue is %+v\n", files, files)

	for _, file := range files {
		fmt.Printf("file type is %T ", file)
		fmt.Println(file.Name())
	}
}
