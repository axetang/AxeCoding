/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.SectionReader
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type SectionReader struct {
    // contains filtered or unexported fields
	r     ReaderAt
	base  int64
	off   int64
	limit int64
 }
 func (s *SectionReader) Read(p []byte) (n int, err error)
 func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
 func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
 func (s *SectionReader) Size() int64
 ------------------------------------------------------------------------------------
 **Description:
 SectionReader implements Read, Seek, and ReadAt on a section of an underlying
 ReaderAt.
 Size returns the size of the section in bytes.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SectionReader结构体实现了Reader、ReaderAt、Seeker三个接口；
 2. SectionReader有一个内部私有成员r，类型为ReaderAt接口；
 3. Size方法返回SectionReader结构体中Reader的len，具体实现需要看源代码确认；
 4. Copy函数调用后，当前reader的读取位置已经到尾部，继续读取将EOF，可以通过Seek来调整位置重新读取；
 *************************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 50)
	s.Seek(0, io.SeekStart)
	n, err := s.Read(buf)
	fmt.Println("\nbuf is ", string(buf), n, "bytes read.", err)

	buf1 := make([]byte, 6)
	s.Seek(0, io.SeekStart)
	if _, err := s.ReadAt(buf1, 8); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf1)
	fmt.Println("size of s : ", s.Size())
}
