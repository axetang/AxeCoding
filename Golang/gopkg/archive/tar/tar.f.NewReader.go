/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.NewReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewReader(r io.Reader) *Reader
 ------------------------------------------------------------------------------------
 **Description:
 NewReader creates a new Reader reading from r.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewReader函数创建一个从r读取的Reader。
*************************************************************************************/
package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names: George-Geoffrey-Gonzo"},
		{"todo.txt", "Get animal handling license."},
		{"Axe Header", "Axe's File body"},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
