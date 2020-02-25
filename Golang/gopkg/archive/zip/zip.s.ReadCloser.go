/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.ReadCloser
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ReadCloser struct {
     Reader
     // contains filtered or unexported fields
 }
 func (rc *ReadCloser) Close() error
 ------------------------------------------------------------------------------------
 **Description:
 Close closes the Zip file, rendering it unusable for I/O.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. ReadCloser结构体显式封装了Reader结构体；
 1. Close方法关闭Zip文件，不再可以继续IO操作。
*************************************************************************************/
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a buffer to write our archive to.
	//buf := new(bytes.Buffer)
	fi, _ := os.Create("./_iofiles/axe.zip")
	// Create a new zip archive.
	w := zip.NewWriter(fi)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names: GeorgeGeoffreyGonzo"},
		{"todo.txt", "Get animal handling licence. Write more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(buf)
	r, _ := zip.OpenReader("./_iofiles/axe.zip")
	for _, f := range r.File {
		rc, _ := f.Open()
		fmt.Println(f.FileHeader.Name)
		io.CopyN(os.Stdout, rc, 200)
		fmt.Println()
	}

}
