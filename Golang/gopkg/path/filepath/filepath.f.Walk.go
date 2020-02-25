/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Walk
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Walk(root string, walkFn WalkFunc) error
 -----------------------------------------------------------------------------------
 **Description:
 Walk walks the file tree rooted at root, calling walkFn for each file or
 directory in the tree, including root. All errors that arise visiting files
 and directories are filtered by walkFn. The files are walked in lexical order,
 which makes the output deterministic but means that for very large directories
 Walk can be inefficient. Walk does not follow symbolic links.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 使用walkFunc遍历处理path目录下的所的文件，包括path；
 2. 参数说明：
	- root: 需要遍历的目录
	- walkFunc: 处理单个文件的函数(linux下目录和文件统称为文件),类型为WalkFunc, WalkFunc
		的定义为type WalkFunc func(path string, info os.FileInfo, err error) error。
		path为文件名，info文件的信息，遍历中的错误。
 3. 返回值：error
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
