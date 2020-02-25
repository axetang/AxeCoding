/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsExist(err error) bool
 func IsNotExist(err error) bool
 func IsPathSeparator(c uint8) bool
 func IsPermission(err error) bool
 func IsTimeout(err error) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsExist returns a boolean indicating whether the error is known to report that a
 file or directory already exists. It is satisfied by ErrExist as well as some
 syscall errors.
 IsNotExist returns a boolean indicating whether the error is known to report that
 a file or directory does not exist. It is satisfied by ErrNotExist as well as some
 syscall errors.
 IsPathSeparator reports whether c is a directory separator character.
 IsPermission returns a boolean indicating whether the error is known to report that
 permission is denied. It is satisfied by ErrPermission as well as some syscall errors.
 IsTimeout returns a boolean indicating whether the error is known to report that a
 timeout occurred.
  ------------------------------------------------------------------------------------
 **要点总结:
 1. IsExist函数判断一个错误err error是否为文件或目录已存在的错误;
 2. IsNotExist函数判断一个错误err error是否为文件或目录不存在的错误;
 3. IsPathSeparator函数判断一个字符是否为目录路径分割符；
 4. IsPermission函数判断一个错误err error是否为无权限错误；
 5. IsTimeout函数判断是否一个错误err error是否为超时出现的错误。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("-----IsExist-----")
	fmt.Printf("%t\n", os.IsExist(os.ErrExist))
	fmt.Printf("%t\n", os.IsExist(errors.New("Custom Error")))

	fmt.Println("-----IsNotExist-----")
	fmt.Printf("%t\n", os.IsNotExist(os.ErrNotExist))
	fmt.Printf("%t\n", os.IsNotExist(errors.New("Custom Error")))

	fmt.Println("-----IsPathSeparator-----")
	fmt.Printf("%t\n", os.IsPathSeparator('/'))
	fmt.Printf("%t\n", os.IsPathSeparator('.'))

	fmt.Println("-----IsPermissionr-----")
	fmt.Printf("%t\n", os.IsPermission(os.ErrPermission))
	fmt.Printf("%t\n", os.IsPermission(errors.New("Custom Error")))

	fmt.Println("-----IsTimeout-----")
	//fmt.Printf("%t\n", os.IsTimeout(os.ErrTimeout))
	fmt.Printf("%t\n", os.IsTimeout(errors.New("Custom Error")))

}
