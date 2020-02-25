/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Chdir
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Chdir(dir string) error
  ------------------------------------------------------------------------------------
 **Description:
 Chdir changes the current working directory to the named directory. If there is an
 error, it will be of type *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 更改当前工作目录到参数dir指定的目录；
 2. 如果出错，返回*PathError。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The current directory is: %s\n", pwd)

	if err = os.Chdir("./_iofiles"); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	pwd, err = os.Getwd()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now the current directory is: %s\n", pwd)
}
