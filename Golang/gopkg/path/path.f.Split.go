/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Split
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Split(path string) (dir, file string)
 ------------------------------------------------------------------------------------
 **Description:
 Split splits path immediately following the final slash, separating it into a
 directory and file name component. If there is no slash in path, Split returns
 an empty dir and file set to path. The returned values have the property that
 path = dir+file.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Split函数分离path参数，分离成为一个目录和一个文件名；
 2. 如果path中没有"\"，Split函数返回一个空目录以及文件；
 3. 返回的值表达的是path=dir+file
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Split("static/myfile.css"))  // static/ myfile.css
	fmt.Println(path.Split("static/myfile.css/")) // static/ myfile.css
	fmt.Println(path.Split("static"))             // "" static
}
