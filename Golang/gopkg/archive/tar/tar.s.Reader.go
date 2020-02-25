/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.Reader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // Reader provides sequential access to the contents of a tar archive.
 // Reader.Next advances to the next file in the archive (including the first),
 // and then Reader can be treated as an io.Reader to access the file's data.
 type Reader struct {
    // contains filtered or unexported fields
	r    io.Reader
	pad  int64      // Amount of padding (ignored) after current file entry
	curr fileReader // Reader for current file entry
	blk  block      // Buffer to use as temporary local storage

	// err is a persistent error.
	// It is only the responsibility of every exported method of Reader to
	// ensure that this error is sticky.
	err error
 }
 func (tr *Reader) Next() (*Header, error)
 func (tr *Reader) Read(b []byte) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 Reader provides sequential access to the contents of a tar archive. Reader.Next
 advances to the next file in the archive (including the first), and then Reader can
 be treated as an io.Reader to access the file's data.
 NewReader creates a new Reader reading from r.
 Next advances to the next entry in the tar archive. The Header.Size determines how
 many bytes can be read for the next file. Any remaining data in the current file is
 automatically discarded.
 io.EOF is returned at the end of the input.
 Read reads from the current file in the tar archive. It returns (0, io.EOF) when it
 reaches the end of that file, until Next is called to advance to the next file.
 If the current file is sparse, then the regions marked as a hole are read back as
 NUL-bytes.
 Calling Read on special types like TypeLink, TypeSymlink, TypeChar, TypeBlock,
 TypeDir, and TypeFifo returns (0, io.EOF) regardless of what the Header.Size claims.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Reader提供了对一个tar档案文件的顺序读取。一个tar档案文件包含一系列文件；
 1. Next方法返回档案中的下一个文件（包括第一个），返回值可以被视为io.Reader来获取文件的数据；Next
 转入tar档案文件下一记录，它会返回下一记录的头域；
 2. Read方法从档案文件的当前记录读取数据到参数b []byte，到达记录末端时返回(0, EOF)，
 直到调用Next方法转入下一记录。
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
