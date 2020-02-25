/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.WalkFunc
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type WalkFunc func(path string, info os.FileInfo, err error) error
 -----------------------------------------------------------------------------------
 **Description:
 WalkFunc is the type of the function called for each file or directory visited
 by Walk. The path argument contains the argument to Walk as a prefix; that is,
 if Walk is called with "dir", which is a directory containing the file "a",
 the walk function will be called with argument "dir/a". The info argument is
 the os.FileInfo for the named path.
 If there was a problem walking to the file or directory named by path, the
 incoming error will describe the problem and the function can decide how to
 handle that error (and Walk will not descend into that directory). In the
 case of an error, the info argument will be nil. If an error is returned,
 processing stops. The sole exception is when the function returns the special
 value SkipDir. If the function returns SkipDir when invoked on a directory,
 Walk skips the directory's contents entirely. If the function returns SkipDir
 when invoked on a non-directory file, Walk skips the remaining files in the
 containing directory.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s %s %t\n", path, info.Name(), info.IsDir())
	return nil
}

func main() {
	//遍历打印所有的文件名
	filepath.Walk("/Users/axe/Downloads", walkFunc)
}
