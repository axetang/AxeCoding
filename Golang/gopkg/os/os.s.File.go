/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.File
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type File struct {
     // contains filtered or unexported fields
 }
 func Create(name string) (*File, error)
 func NewFile(fd uintptr, name string) *File
 func Open(name string) (*File, error)
 func OpenFile(name string, flag int, perm FileMode) (*File, error)
 func (f *File) Chdir() error
 func (f *File) Chmod(mode FileMode) error
 func (f *File) Chown(uid, gid int) error
 func (f *File) Close() error
 func (f *File) Fd() uintptr
 func (f *File) Name() string
 func (f *File) Read(b []byte) (n int, err error)
 func (f *File) ReadAt(b []byte, off int64) (n int, err error)
 func (f *File) Readdir(n int) ([]FileInfo, error)
 func (f *File) Readdirnames(n int) (names []string, err error)
 func (f *File) Seek(offset int64, whence int) (ret int64, err error)
 func (f *File) SetDeadline(t time.Time) error
 func (f *File) SetReadDeadline(t time.Time) error
 func (f *File) SetWriteDeadline(t time.Time) error
 func (f *File) Stat() (FileInfo, error)
 func (f *File) Sync() error
 func (f *File) Truncate(size int64) error
 func (f *File) Write(b []byte) (n int, err error)
 func (f *File) WriteAt(b []byte, off int64) (n int, err error)
 func (f *File) WriteString(s string) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 File represents an open file descriptor.
 Create creates the named file with mode 0666 (before umask), truncating it if it
 already exists. If successful, methods on the returned File can be used for I/O; the
 associated file descriptor has mode O_RDWR. If there is an error, it will be of type
 *PathError.
 NewFile returns a new File with the given file descriptor and name. The returned
 value will be nil if fd is not a valid file descriptor. On Unix systems, if the file
 descriptor is in non-blocking mode, NewFile will attempt to return a pollable File
 (one for which the SetDeadline methods work).
 Open opens the named file for reading. If successful, methods on the returned file
 can be used for reading; the associated file descriptor has mode O_RDONLY. If there
 is an error, it will be of type *PathError.
 OpenFile is the generalized open call; most users will use Open or Create instead. It
 opens the named file with specified flag (O_RDONLY etc.) and perm (before umask), if
 applicable. If successful, methods on the returned File can be used for I/O. If there
 is an error, it will be of type *PathError.
 Chdir changes the current working directory to the file, which must be a directory.
 If there is an error, it will be of type *PathError.
 Chmod changes the mode of the file to mode. If there is an error, it will be of type
 *PathError.
 Chown changes the numeric uid and gid of the named file. If there is an error, it
 will be of type *PathError. On Windows, it always returns the syscall.EWINDOWS error,
 wrapped in *PathError.
 Close closes the File, rendering it unusable for I/O. On files that support
 SetDeadline, any pending I/O operations will be canceled and return immediately with
 an error.
 Fd returns the integer Unix file descriptor referencing the open file.
 The file descriptor is valid only until f.Close is called or f is garbage collected.
 On Unix systems this will cause the SetDeadline methods to stop working.
 Name returns the name of the file as presented to Open.
 Read reads up to len(b) bytes from the File. It returns the number of bytes read and
 any error encountered. At end of file, Read returns 0, io.EOF.
 ReadAt reads len(b) bytes from the File starting at byte offset off. It returns the
 number of bytes read and the error, if any. ReadAt always returns a non-nil error
 when n < len(b). At end of file, that error is io.EOF.
 Readdir reads the contents of the directory associated with file and returns a slice
 of up to n FileInfo values, as would be returned by Lstat, in directory order.
 Subsequent calls on the same file will yield further FileInfos.
    If n > 0, Readdir returns at most n FileInfo structures. In this case, if Readdir
    returns an empty slice, it will return a non-nil error explaining why. At the end
    of a directory, the error is io.EOF.
    If n <= 0, Readdir returns all the FileInfo from the directory in a single slice.
    In this case, if Readdir succeeds (reads all the way to the end of the directory),
    it returns the slice and a nil error. If it encounters an error before the end of
    the directory, Readdir returns the FileInfo read until that point and a non-nil
    error.
 Readdirnames reads and returns a slice of names from the directory f.
    If n > 0, Readdirnames returns at most n names. In this case, if Readdirnames
    returns an empty slice, it will return a non-nil error explaining why. At the
    end of a directory, the error is io.EOF.
    If n <= 0, Readdirnames returns all the names from the directory in a single slice.
    In this case, if Readdirnames succeeds (reads all the way to the end of the
    directory), it returns the slice and a nil error. If it encounters an error before
    the end of the directory, Readdirnames returns the names read until that point and
    a non-nil error.
 Seek sets the offset for the next Read or Write on file to offset, interpreted
 according to whence: 0 means relative to the origin of the file, 1 means relative to
 the current offset, and 2 means relative to the end. It returns the new offset and an
 error, if any. The behavior of Seek on a file opened with O_APPEND is not specified.
 SetDeadline sets the read and write deadlines for a File. It is equivalent to calling
 both SetReadDeadline and SetWriteDeadline.
 Only some kinds of files support setting a deadline. Calls to SetDeadline for files
 that do not support deadlines will return ErrNoDeadline. On most systems ordinary
 files do not support deadlines, but pipes do.
 A deadline is an absolute time after which I/O operations fail with an error instead
 of blocking. The deadline applies to all future and pending I/O, not just the
 immediately following call to Read or Write. After a deadline has been exceeded,
 the connection can be refreshed by setting a deadline in the future.
 An error returned after a timeout fails will implement the Timeout method, and
 calling the Timeout method will return true. The PathError and SyscallError types
 implement the Timeout method. In general, call IsTimeout to test whether an error
 indicates a timeout.
 An idle timeout can be implemented by repeatedly extending the deadline after
 successful Read or Write calls.
 A zero value for t means I/O operations will not time out.
 SetReadDeadline sets the deadline for future Read calls and any currently-blocked
 Read call. A zero value for t means Read will not time out. Not all files support
 setting deadlines; see SetDeadline.
 SetWriteDeadline sets the deadline for any future Write calls and any
 currently-blocked Write call. Even if Write times out, it may return n > 0,
 indicating that some of the data was successfully written. A zero value for t means
 Write will not time out. Not all files support setting deadlines; see SetDeadline.
 Stat returns the FileInfo structure describing file. If there is an error, it will
 be of type *PathError.
 Sync commits the current contents of the file to stable storage. Typically, this
 means flushing the file system's in-memory copy of recently written data to disk.
 Truncate changes the size of the file. It does not change the I/O offset. If there is
 an error, it will be of type *PathError.
 Write writes len(b) bytes to the File. It returns the number of bytes written and an
 error, if any. Write returns a non-nil error when n != len(b).
 WriteAt writes len(b) bytes to the File starting at byte offset off. It returns the
 number of bytes written and an error, if any. WriteAt returns a non-nil error when
 n != len(b).
 WriteString is like Write, but writes the contents of string s rather than a slice
 of bytes.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 结构体File代表一个打开文件的描述符；
 1. 查看源码，File结构体定义如下，包含一个匿名的file结构体指针。继续查看file结构体定义，具体如下；
 type File struct {
	*file // os specific
 }
 type file struct {
	pfd         poll.FD
	name        string
	dirinfo     *dirInfo // nil unless directory being read
	nonblock    bool     // whether we set nonblocking mode
	stdoutOrErr bool     // whether this is stdout or stderr
 }
 2. Create方法创建一个文件，此方法会覆盖已存在的文件;
 3. NewFile方法新建一个文件（但不保存），返回文件指针；
 4. Fd方法获取和返回对应的文件句柄uintptr类型；
 5. Open方法以只读的模式打开一个文件；
 6. OpenFile方法打开一个文件，打开的方式和权限通过参数flag和perm控制；
 7. ChDir方法改变当前工作目录，参数f *File必须是一个目录的信息；
 8. Chmod方法修改文件的mode；
 9. Chown方法修改文件所属的用户和组；
 10. Close方法关闭文件；
 11. Name方法返回文件名；
 12. Read方法读取文件到参数b []byte中，最多读取len(b)长度的数据；
 13. ReadAt方法从参数off指定位置处读取文件到参数b []bytes中，最多读取len(b)个字节；
 14. ReadDir方法读取目录下的文件信息，如果参数n>0，则读取min(n, 文件夹下的文件数目)个文件信息，
 如果n<=0，则读取文件夹下所有的文件信息；文件信息以切片[]FileInfo返回；
 15. ReadDirName方法读取目录下的文件名称信息，如果参数n>0，则读取min(n, 文件夹下的文件数目)个文件
 如果n<=0，则读取文件夹下所有的文件；文件名称以切片[]string返回；
 16. Seek方法把文件指针移动到文件的指定位置的offset相对偏移量，whence为0时代表相对文件开始的位置，
 1代表相对当前位置，2代表相对文件结尾的位；
 17. SetDeadline：后续补充
 18. SetReadDeadline：后续补充
 19. SetWriteDeadline：后续补充
 20. Stat方法获取文件信息，并返回包含文件信息的FileInfo接口实例；
 21. Sync方法把f内存里的内容写到磁盘上，即使是断电或系统崩溃，也能做到数据不丢失（注：一般不用调用
 这个方法）；
 22. Truncate把文件截断到参数size int16指定的大小；
 23. Write方法向文件写入b []byte，最多写入len(b)字节内容；
 24. WriteAt在指定偏移量off处写入len(b)字节内容；
 25. WriteString类似Write，不过写入文件的是参数字符串s。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	/*fmt.Println("--------Create&Name----------")
			TestCreate()
			fmt.Println("--------NewFile----------")
			TestNewFile()
			fmt.Println("--------Fd----------")
			TestFd()
			fmt.Println("--------Open----------")
			TestOpen()
			fmt.Println("--------OpenFile----------")
			TestOpenFile()
			fmt.Println("--------Chdir----------")
			TestChdir()
			fmt.Println("--------Chmod----------")
			TestChmod()
			fmt.Println("--------Chown----------")
			TestChown()
			fmt.Println("--------Close----------")
	        TestClose()
	        fmt.Println("--------Name----------")
			TestName()
			fmt.Println("--------Read----------")
			TestRead()
			fmt.Println("--------ReadAt----------")
			TestReadAt()
			fmt.Println("--------Readdir----------")
			TestReaddir()
			fmt.Println("--------Readdirnames----------")
		    TestReaddirnames()
		fmt.Println("--------Seek----------")
		TestSeek()
		fmt.Println("--------Stat----------")
		TestStat()
		fmt.Println("--------Sync----------")
		TestSync()
	fmt.Println("--------Truncate----------")
	TestTruncate()
	fmt.Println("--------Write----------")
	TestWrite()*/
	fmt.Println("--------WriteAt----------")
	//TestWriteAt()
	fmt.Println("--------WriteString----------")
	TestWriteString()

}

func TestWrite() {
	b := make([]byte, 10)
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR|os.O_APPEND, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	n, _ := fi.Read(b)
	fmt.Printf("len: %d, file content: %s\n", n, b[:n])
	fi.Write([]byte("cc\n"))
	fi.Seek(0, 0)
	n, _ = fi.Read(b)
	fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
}

func TestWriteAt() {
	b := make([]byte, 10)
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	n, _ := fi.Read(b)
	fmt.Printf("len: %d, file content: %s\n", n, b[:n])
	fi.Seek(0, 0)
	fi.WriteAt([]byte("cc\n"), 3)
	fi.Seek(0, 0)
	n, _ = fi.Read(b)
	fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
}

func TestWriteString() {
	b := make([]byte, 20)
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR|os.O_APPEND, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	n, _ := fi.Read(b)
	fmt.Printf("len: %d, file content: %s\n", n, b[:n])
	fi.WriteString("Hello World!\n")
	fi.Seek(0, 0)
	n, _ = fi.Read(b)
	fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
}

func TestSync() {
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR|os.O_APPEND, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fi.Write([]byte("cc\n"))
	fi.Sync()
}

func TestTruncate() {
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	f, _ := fi.Stat()
	fmt.Printf("file size: %+v\n", f.Size())
	f, _ = fi.Stat()
	fi.Truncate(f.Size() - 2)
	fmt.Printf("now file size: %+v\n", f.Size())
}

func TestStat() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	f, _ := fi.Stat()
	fmt.Printf("file info: %+v\n", f)
}

func TestSeek() {
	b := make([]byte, 10)
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fi.Seek(3, 0)
	n, _ := fi.Read(b)
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", b[:n])
}

func TestReaddirnames() {
	fi, err := os.Open("./_iofiles/")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fn, _ := fi.Readdirnames(10)
	for i, f := range fn {
		fmt.Printf("fileinfo %d: %+v\n", i, f)
	}
}
func TestReaddir() {
	fi, err := os.Open("./_iofiles/")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	ff, _ := fi.Readdir(10)
	for i, f := range ff {
		fmt.Printf("fileinfo %d: %+v\n", i, f)
	}
}

func TestReadAt() {
	b := make([]byte, 10)
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	n, _ := fi.ReadAt(b, 3)
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", b[:n])
}

func TestRead() {
	b := make([]byte, 10)
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	n, _ := fi.Read(b)
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", b[:n])
}

func TestName() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("The file name is: %v\n", fi.Name())
}

func TestChown() {
	fi, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)

	f, _ := os.Open("./_iofiles/Hello.go")
	err = f.Chown(99, 99) //nobody, nobody
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	f.Close()

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)
}

func TestChmod() {
	fi, err := os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)

	f, _ := os.Open("./_iofiles/Hello.go")
	err = f.Chmod(0777)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	f.Close()

	fi, err = os.Stat("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)
}

func TestChdir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The current directory is: %s\n", pwd)

	f, _ := os.Open("_i1ofiles")
	if err = f.Chdir(); err != nil {
		fmt.Printf("Error1: %s,%s\n", err)
	}

	pwd, err = os.Getwd()
	if err != nil {
		fmt.Printf("Error2: %s\n", err)
		return
	}
	fmt.Printf("Now the current directory is: %s\n", pwd)
}

func TestOpen() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	data := make([]byte, 100)
	fi.Read(data)
	fmt.Printf("The %s's content is: \n%s \n", fi.Name(), data)
}

func TestOpenFile() {
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	data := make([]byte, 100)
	fi.Read(data)
	fmt.Printf("The %s's content is: \n%s \n", fi.Name(), data)
	fi.WriteString("come on!!\n")
	fi.Seek(0, 0)
	fi.Read(data)
	fmt.Printf("Now the %s's content is:\n%s \n", fi.Name(), data)
}

func TestFd() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("type：%T, Value: %v\n", fi.Fd(), fi.Fd())
}

func TestNewFile() {
	fi, err := os.Open("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("type：%T, Value: %v\n", fi.Fd(), fi.Fd())
	file := os.NewFile(fi.Fd(), "./_iofiles/World.go")
	defer file.Close()
	fmt.Printf("The %s has been new!\n", file.Name())
}

func TestCreate() {
	fi, err := os.Create("./_iofiles/Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	fmt.Printf("The %s has been created!\n", fi.Name())
}
