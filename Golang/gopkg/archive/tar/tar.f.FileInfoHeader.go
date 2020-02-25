/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.FileInfoHeader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
 ------------------------------------------------------------------------------------
 **Description:
 FileInfoHeader creates a partially-populated Header from fi. If fi describes a
 symlink, FileInfoHeader records link as the link target. If fi describes a
 directory, a slash is appended to the name.
 Since os.FileInfo's Name method only returns the base name of the file it
 describes, it may be necessary to modify Header.Name to provide the full path
 name of the file.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FileInfoHeader函数返回一个根据fi填写了部分字段的Header。 如果fi描述一个符号链接，
 FileInfoHeader函数将link参数作为链接目标。如果fi描述一个目录，会在名字后面添加斜杠。
 因为os.FileInfo接口的Name方法只返回它描述的文件的无路径名，有可能需要将返回值的Name字段修改为
 文件的完整路径名。
*************************************************************************************/
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	//var buf bytes.Buffer
	var tf1, _ = os.Open("./_iofiles/tarfile1.tar")
	tf2, _ := os.Create("./_iofiles/tarfile2.tar")
	defer tf1.Close()
	fi, _ := os.Stat("./_iofiles/tarfile1.tar")
	fih, _ := tar.FileInfoHeader(fi, "")
	tw := tar.NewWriter(tf2)
	/*var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names: George-Geoffrey-Gonzo"},
		{"todo.txt", "Get animal handling license."},
		{"Axe Header", "Axe's File body"},
	}*/

	tw.WriteHeader(fih)
	tw.Write([]byte("I am Axe!\n"))

	// Open and iterate through the files in the archive.
	tf3, _ := os.Open("./_iofiles/tarfile2.tar")
	defer tf3.Close()
	tr := tar.NewReader(tf3)
	for {
		hdr, err := tr.Next()
		fmt.Println("enter newreader for loop")
		if err == io.EOF {
			fmt.Println("Next() EOF")
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
			fmt.Println("Fatal err", err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			fmt.Println("error in Copy")
			log.Fatal(err)
		}
		fmt.Println()
	}
}
