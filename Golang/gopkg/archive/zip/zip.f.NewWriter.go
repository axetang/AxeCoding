/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.NewWriter
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewWriter(w io.Writer) *Writer
 ------------------------------------------------------------------------------------
 **Description:
 NewWriter returns a new Writer writing a zip file to w.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewWriter创建并返回一个将zip文件写入w的*Writer。
*************************************************************************************/
package main

import (
	"archive/zip"
	"bytes"
	"log"
)

func main() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
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
}
