/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ioutil
 **Element: ioutil.WriteFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func WriteFile(filename string, data []byte, perm os.FileMode) error
 ------------------------------------------------------------------------------------
 **Description:
 WriteFile writes data to a file named by filename. If the file does not exist,
 WriteFile creates it with permissions perm; otherwise WriteFile truncates it
 before writing.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. WriteFile将数据写入到名为filename的文件中;
 2. 若该文件不存在，WriteFile就会按照权限perm创建它；
 3. 若该文件存在，WriteFile就会按照覆盖重写该文件。
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("./_iofiles", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir is ", dir)

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tmpfn type is %T, value is %v\n", tmpfn, tmpfn)

	var i int
	fmt.Scanf("\n%d", i)

	//Write again with the existed filename
	content1 := []byte("New Content for WriteFile again@!")
	if err := ioutil.WriteFile(tmpfn, content1, 0666); err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Scanf("\n%d", i)
}
