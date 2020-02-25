/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.Reader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Reader struct {
     File    []*File
     Comment string
     // contains filtered or unexported fields
 }
 func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
 ------------------------------------------------------------------------------------
 **Description:
 RegisterDecompressor registers or overrides a custom decompressor for a specific
 method ID. If a decompressor for a given method is not found, Reader will default
 to looking up the decompressor at the package level.
 ------------------------------------------------------------------------------------
 **要点总结:
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
	r, err := zip.OpenReader("testdata/readme.zip")
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
