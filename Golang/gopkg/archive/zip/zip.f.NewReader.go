/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.NewReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReader(r io.ReaderAt, size int64) (*Reader, error)
 ------------------------------------------------------------------------------------
 **Description:
 NewReader returns a new Reader reading from r, which is assumed to have the given
 size in bytes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewReader返回一个从r读取数据的*Reader，r被假设其大小为size字节。
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
	r, err := zip.OpenReader("./_iofiles/pkgzip.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
