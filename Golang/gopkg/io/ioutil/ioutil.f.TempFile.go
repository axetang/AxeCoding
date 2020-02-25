/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ioutil
 **Element: ioutil.TempFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func TempFile(dir, pattern string) (f *os.File, err error)
 ------------------------------------------------------------------------------------
 **Description:
 TempFile creates a new temporary file in the directory dir, opens the file for
 reading and writing, and returns the resulting *os.File. The filename is generated
 by taking pattern and adding a random string to the end. If pattern includes a "*",
 the random string replaces the last "*". If dir is the empty string, TempFile uses
 the default directory for temporary files (see os.TempDir). Multiple programs
 calling TempFile simultaneously will not choose the same file. The caller can use
 f.Name() to find the pathname of the file. It is the caller's responsibility to
 remove the file when no longer needed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. TempFile在目录dir中创建一个名字以prefix开头+随机数的新的临时文件，打开该文件以用于读写，执行
 完毕返回*os.File；注意要调用该文件的Close方法关闭文件句柄；
 2. 若dir为空字符串，TempFile就会为临时文件使用默认的目录（见os.TempDir）;
 3. 多程序同时调用TempFile将不会选择相同的文件，因为文件名是prefix+随机数；
 4. 调用者可使用 f.Name()来查找该文件的路径名pathname；
 5. 当该文件不再被需要时，调用者应负责将其移除，可采用defer函数来处理;
 6. prefix中的字符串如果包含通配符*，则通配符位置使用随机数占据，如下例子程序。
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("./_iofiles", "Example.*.txt")
	if err != nil E
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	var i int
	fmt.Scanf("%d", i)
}
