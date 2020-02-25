/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.Writer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // Writer provides sequential writing of a tar archive.
 // Write.WriteHeader begins a new file with the provided Header,
 // and then Writer can be treated as an io.Writer to supply that file's data.
 type Writer struct {
    // contains filtered or unexported fields
	w    io.Writer
	pad  int64      // Amount of padding to write after current file entry
	curr fileWriter // Writer for current file entry
	hdr  Header     // Shallow copy of Header that is safe for mutations
	blk  block      // Buffer to use as temporary local storage

	// err is a persistent error.
	// It is only the responsibility of every exported method of Writer to
	// ensure that this error is sticky.
	err error
 }
 func (tw *Writer) Close() error
 func (tw *Writer) Flush() error
 func (tw *Writer) Write(b []byte) (int, error)
 func (tw *Writer) WriteHeader(hdr *Header) error
 ------------------------------------------------------------------------------------
 **Description:
 Writer provides sequential writing of a tar archive. Write.WriteHeader begins a new
 file with the provided Header, and then Writer can be treated as an io.Writer to
 supply that file's data.
 Close closes the tar archive by flushing the padding, and writing the footer. If
 the current file (from a prior call to WriteHeader) is not fully written, then
 this returns an error.
 Flush finishes writing the current file's block padding. The current file must be
 fully written before Flush can be called.
 This is unnecessary as the next call to WriteHeader or Close will implicitly flush
 out the file's padding.
 Write writes to the current file in the tar archive. Write returns the error
 ErrWriteTooLong if more than Header.Size bytes are written after WriteHeader.
 Calling Write on special types like TypeLink, TypeSymlink, TypeChar, TypeBlock,
 TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless of what the
 Header.Size claims.
 WriteHeader writes hdr and prepares to accept the file's contents. The Header.Size
 determines how many bytes can be written for the next file. If the current file
 is not fully written, then this returns an error. This implicitly flushes any
 padding necessary before writing the header.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Writer类型提供了POSIX.1格式的tar档案文件的顺序写入。一个tar档案文件包含一系列文件。调用
 WriteHeader来写入一个新的文件，然后调用Write写入文件的数据，该记录写入的数据不能超过hdr.Size字节；
 1. WriteHeader方法写入hdr *Header并准备接受文件内容。如果不是第一次调用本方法，会调用Flush;
 在Close之后调用本方法会返回ErrWriteAfterClose；
 2. Write方法向tar档案文件的当前记录中写入文件内容数据。如果写入的数据总数超出上一次调用
 WriteHeader的参数hdr.Size字节，返回ErrWriteTooLong错误；
 3. 以二进制方式写入。
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
