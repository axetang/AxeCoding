/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.FileMode
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type FileMode uint32
 const (
     // The single letters are the abbreviations
     // used by the String method's formatting.
     ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
     ModeAppend                                     // a: append-only
     ModeExclusive                                  // l: exclusive use
     ModeTemporary                                  // T: temporary file; Plan 9 only
     ModeSymlink                                    // L: symbolic link
     ModeDevice                                     // D: device file
     ModeNamedPipe                                  // p: named pipe (FIFO)
     ModeSocket                                     // S: Unix domain socket
     ModeSetuid                                     // u: setuid
     ModeSetgid                                     // g: setgid
     ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
     ModeSticky                                     // t: sticky
     ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

     // Mask for the type bits. For regular files, none will be set.
     ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeIrregular

     ModePerm FileMode = 0777 // Unix permission bits
 )
 The defined file mode bits are the most significant bits of the FileMode. The nine least-significant bits are the standard Unix rwxrwxrwx permissions. The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.

 func (m FileMode) IsDir() bool
 func (m FileMode) IsRegular() bool
 func (m FileMode) Perm() FileMode
 func (m FileMode) String() string
 ------------------------------------------------------------------------------------
 **Description:
 A FileMode represents a file's mode and permission bits. The bits have the same
 definition on all systems, so that information about files can be moved from one
 system to another portably. Not all bits apply to all systems. The only required bit
 is ModeDir for directories.
 IsDir reports whether m describes a directory. That is, it tests for the ModeDir bit
 being set in m.
 IsRegular reports whether m describes a regular file. That is, it tests that no mode
 type bits are set.
 Perm returns the Unix permission bits in m.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FileMode定义为uint32，代表文件状态，可以用字符串“dalTLDpSugct?”代表，具体见常数定义；
 2. File结构体有一个方法Stat，获取并返回file的具体信息即*FileInfo接口指针，实际上指向fileStat，
 fileStat结构体实现了FileInfo接口；
 3. 实现FileInfo接口的结构体fileStat实现了Mode方法，返回FileMode结构体信息；
 4. FileMode结构体的方法Perm返回文件的Perm信息，类型依然是FileMode；而IsDir方法判断FileMode是否
 等于ModeDir；
 5. 基于以上分析，因此有了如下程序中的表达：
    fi,_ := f.Stat()
	fmt.Printf("%s is dir?: %t\n", fi.Name(), fi.Mode().IsDir())
    fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().Perm())
	fmt.Printf("%s is regular?: %t\n", fi.Name(), fi.Mode().IsRegular())
	fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().String())
6. String方法返回FileMode Perm的字符串表达式，具体实现可参看标准库源码。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(uint32(os.ModeDir), uint32(os.ModeAppend))
	b := []byte("dalTLDpSugct?")
	fmt.Printf("b is %b\n", b)
	fmt.Printf("d is %b\n", uint32(os.ModeDir))
	fmt.Printf("a is %b\n", uint32(os.ModeAppend))

	println("--IsDir--")
	TestIsDir()
	println("--Perm--")
	TestPerm()
	println("--IsRegular--")
	TestIsRegular()
	println("--String--")
	TestString()
}

func TestString() {
	f, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()
	fi, _ := f.Stat()
	fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().String())
}

func TestIsRegular() {
	f, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()
	fi, _ := f.Stat()
	fmt.Printf("%s is regular?: %t\n", fi.Name(), fi.Mode().IsRegular())
}

func TestIsDir() {
	f, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()
	fi, _ := f.Stat()
	fmt.Printf("%s is dir?: %t\n", fi.Name(), fi.Mode().IsDir())
}
func TestPerm() {
	f, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()
	fi, _ := f.Stat()
	//fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().Perm()&0777)
	fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().Perm())
}
