/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.LookPath
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookPath(file string) (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 LookPath searches for an executable named file in the directories named by the PATH
 environment variable. If file contains a slash, it is tried directly and the PATH
 is not consulted. The result may be an absolute path or a path relative to the
 current directory.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. LookPath根据环境变量PATH来搜寻可执行文件file的路径；
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("ls is available at %s\n", path)
}
