/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.OpenReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func OpenReader(name string) (*ReadCloser, error)
 ------------------------------------------------------------------------------------
 **Description:
 OpenReader will open the Zip file specified by name and return a ReadCloser.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. OpenReader会打开参数name string指定的zip文件并返回一个*ReadCloser；
 2. ReadCloser结构体封装了Reader结构体，Reader结构体包含File []*File的成员，存储着zip里的文件信息。
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
	// Open a zip archive for reading.
	r, err := zip.OpenReader("./_iofiles/Archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Iterate through the files in the archive,
	// printing some of their contents.
	fmt.Println(r.File)
	fmt.Println("Comment:", r.Comment)
	fmt.Println("r.File[0].Name", r.File[0].Name)
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 192)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
