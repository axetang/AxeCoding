/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ioutil
 **Element: ioutil.ReadFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadFile(filename string) ([]byte, error)
 ------------------------------------------------------------------------------------
 **Description:
 ReadFile reads the file named by filename and returns the contents. A successful
 call returns err == nil, not err == EOF. Because ReadFile reads the whole file, it
 does not treat an EOF from Read as an error to be reported.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ReadFile读取名为filename的文件并以[]byte来返回其字节内容;
 2. 一次成功的调用应当返回err == nil，而非err == EOF。因为ReadFile会读取整个文件，它并不会将来
 自Read的EOF视作错误来报告。
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./ioutil.f.ReadFile.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s\n", content)
}
