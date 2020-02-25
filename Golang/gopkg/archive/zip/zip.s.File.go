/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.File
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type File struct {
     FileHeader
     // contains filtered or unexported fields
 }
 func (f *File) DataOffset() (offset int64, err error)
 func (f *File) Open() (io.ReadCloser, error)
 ------------------------------------------------------------------------------------
 **Description:
 DataOffset returns the offset of the file's possibly-compressed data, relative to
 the beginning of the zip file.
 Most callers should instead use Open, which transparently decompresses data and
 verifies checksums.
 Open returns a ReadCloser that provides access to the File's contents. Multiple
 files may be read concurrently.
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
