/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.Writer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Writer struct {
     // contains filtered or unexported fields
 }
 func (w *Writer) Close() error
 func (w *Writer) Create(name string) (io.Writer, error)
 func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
 func (w *Writer) Flush() error
 func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
 func (w *Writer) SetComment(comment string) error
 func (w *Writer) SetOffset(n int64)
 ------------------------------------------------------------------------------------
 **Description:
 Writer implements a zip file writer.
 Close finishes writing the zip file by writing the central directory. It does not
 close the underlying writer.
 Create adds a file to the zip file using the provided name. It returns a Writer to
 which the file contents should be written. The file contents will be compressed
 using the Deflate method. The name must be a relative path: it must not start with
 a drive letter (e.g. C:) or leading slash, and only forward slashes are allowed.
 To create a directory instead of a file, add a trailing slash to the name. The
 file's contents must be written to the io.Writer before the next call to Create,
 CreateHeader, or Close.
 CreateHeader adds a file to the zip archive using the provided FileHeader for the
 file metadata. Writer takes ownership of fh and may mutate its fields. The caller
 must not modify fh after calling CreateHeader.
 This returns a Writer to which the file contents should be written. The file's
 contents must be written to the io.Writer before the next call to Create,
 CreateHeader, or Close.
 Flush flushes any buffered data to the underlying writer. Calling Flush is not
 normally necessary; calling Close is sufficient.
 RegisterCompressor registers or overrides a custom compressor for a specific method
 ID. If a compressor for a given method is not found, Writer will default to looking
 up the compressor at the package level.
 SetComment sets the end-of-central-directory comment field. It can only be called
 before Close.
 SetOffset sets the offset of the beginning of the zip data within the underlying
 writer. It should be used when the zip data is appended to an existing file, such
 as a binary executable. It must be called before any data is written.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Writer类型实现了zip文件的写入器。
 1. CreateHeader方法使用给出的*FileHeader来作为文件的元数据添加一个文件进zip文件。
 本方法返回一个io.Writer接口（用于写入新添加文件的内容）。新增文件的内容必须在下一次调用
 CreateHeader、Create或Close方法之前全部写入。
 2. Create方法使用给出的文件名添加一个文件进zip文件。本方法返回一个io.Writer接口（用于写入
 新添加文件的内容）。文件名必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔。新增
 文件的内容必须在下一次调用CreateHeader、Create或Close方法之前全部写入。
 3. Close方法通过写入中央目录关闭该*Writer。本方法不会也没办法关闭下层的io.Writer接口。
*************************************************************************************/
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

func TestRegisterCompressor() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}

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

	fmt.Println("----TestRegisterCompressor----")
	TestRegisterCompressor()
}
