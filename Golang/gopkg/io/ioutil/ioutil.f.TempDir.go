/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package:
 **Element:
 **Type:
 ------------------------------------------------------------------------------------
 **Definition:
 func TempDir(dir, prefix string) (name string, err error)
 ------------------------------------------------------------------------------------
 **Description:
 TempDir creates a new temporary directory in the directory dir with a name beginning
 with prefix and returns the path of the new directory. If dir is the empty string,
 TempDir uses the default directory for temporary files (see os.TempDir). Multiple
 programs calling TempDir simultaneously will not choose the same directory. It is
 the caller's responsibility to remove the directory when no longer needed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. TempDir在目录dir中创建一个名字以prefix开头加上一个随机数的新的临时目录并返回该新目录的路径;
 2. 若dir为空字符串，TempDir就会为临时文件（Unix将目录也视作文件）使用默认的目录（见os.TempDir）；
 3. 多程序同时调用TempDir将不会选择相同的目录；
 4. 当该目录不再被需要时，调用者应负责将其移除，可使用defer函数来完成。
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
}
