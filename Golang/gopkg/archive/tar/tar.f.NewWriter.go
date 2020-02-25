/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.NewWriter
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewWriter(w io.Writer) *Writer
 ------------------------------------------------------------------------------------
 **Description:
 NewWriter creates a new Writer writing to w.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewWriter创建一个写入w的*Writer。
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
	var tf1, _ = os.Create("./_iofiles/tarfile1.tar")
	defer tf1.Close()
	tw := tar.NewWriter(tf1)
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
	tf2, _ := os.Open("./_iofiles/tarfile1.tar")
	defer tf2.Close()
	tr := tar.NewReader(tf2)
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
		fmt.Printf("Contents of %s:", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
