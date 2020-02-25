/*************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Copy()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Copy(dst Writer, src Reader) (written int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Copy copies from src to dst until either EOF is
 reached on src or an error occurs. It returns the number of bytes
 copied and the first error encountered while copying, if any.
 A successful Copy returns err == nil, not err == EOF. Because Copy
 is defined to read from src until EOF, it does not treat an EOF
 from Read as an error to be reported. If src implements the
 WriterTo interface, the copy is implemented by calling
 src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom
 interface, the copy is implemented by calling dst.ReadFrom(src).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 要求src和Dest至少分别实现了io.Reader和io.Writer接口;
 2. Copy()读取src时遇到EOF则停止读，并将读出的字节写入dst。在EOF情况下Copy()返回的err是nil；
 3. os.Open()和os.Create()的文件变量类型都是*os.File，这个类型已经实现了io.Reader和io.Writer
 接口，可以作为Copy()的形参。
 4. Copy函数会将src（Reader）的读取位置置于末尾，继续读取将遇到EOF，需要通过Seek来重新调整读取起点
************************************************************************************/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//srcFile's type is *os.File, which implements the io.reader & io.writer interfaces.
	srcFile, oerr := os.Open("./_iofiles/io.Copy.src.txt")
	fmt.Printf("srcFile type is %T, oerr is %s\n", srcFile, oerr)

	//destFile's type is *os.File, which implements the io.reader & io.writer interfaces.
	destFile1, derr := os.Create("./_iofiles/io.Copy.dst1.txt")
	fmt.Printf("destFile1 type is %T, derr is %s\n", destFile1, derr)

	written, err := io.Copy(destFile1, srcFile)
	if err == nil {
		fmt.Println("Copy Success, total", written, "bytes.")
	}

	destFile2, derr := os.Create("./_iofiles/io.Copy.dst2.txt")
	fmt.Printf("destFile2 type is %T, derr is %s\n", destFile2, derr)
	srcFile.Seek(-1000, io.SeekCurrent)
	written, err = io.Copy(destFile2, srcFile)
	if err == nil {
		fmt.Println("Copy Success, total", written, "bytes.")
	}
}
