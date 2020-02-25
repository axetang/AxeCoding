/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.Header
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Header struct {
     // Typeflag is the type of header entry.
     // The zero value is automatically promoted to either TypeReg or TypeDir
     // depending on the presence of a trailing slash in Name.
     Typeflag byte

     Name     string // Name of file entry
     Linkname string // Target name of link (valid for TypeLink or TypeSymlink)

     Size  int64  // Logical file size in bytes
     Mode  int64  // Permission and mode bits
     Uid   int    // User ID of owner
     Gid   int    // Group ID of owner
     Uname string // User name of owner
     Gname string // Group name of owner

     // If the Format is unspecified, then Writer.WriteHeader rounds ModTime
     // to the nearest second and ignores the AccessTime and ChangeTime fields.
     //
     // To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
     // To use sub-second resolution, specify the Format as PAX.
     ModTime    time.Time // Modification time
     AccessTime time.Time // Access time (requires either PAX or GNU support)
     ChangeTime time.Time // Change time (requires either PAX or GNU support)

     Devmajor int64 // Major device number (valid for TypeChar or TypeBlock)
     Devminor int64 // Minor device number (valid for TypeChar or TypeBlock)

     // Xattrs stores extended attributes as PAX records under the
     // "SCHILY.xattr." namespace.
     //
     // The following are semantically equivalent:
     //  h.Xattrs[key] = value
     //  h.PAXRecords["SCHILY.xattr."+key] = value
     //
     // When Writer.WriteHeader is called, the contents of Xattrs will take
     // precedence over those in PAXRecords.
     //
     // Deprecated: Use PAXRecords instead.
     Xattrs map[string]string

     // PAXRecords is a map of PAX extended header records.
     //
     // User-defined records should have keys of the following form:
     //	VENDOR.keyword
     // Where VENDOR is some namespace in all uppercase, and keyword may
     // not contain the '=' character (e.g., "GOLANG.pkg.version").
     // The key and value should be non-empty UTF-8 strings.
     //
     // When Writer.WriteHeader is called, PAX records derived from the
     // other fields in Header take precedence over PAXRecords.
     PAXRecords map[string]string

     // Format specifies the format of the tar header.
     //
     // This is set by Reader.Next as a best-effort guess at the format.
     // Since the Reader liberally reads some non-compliant files,
     // it is possible for this to be FormatUnknown.
     //
     // If the format is unspecified when Writer.WriteHeader is called,
     // then it uses the first format (in the order of USTAR, PAX, GNU)
     // capable of encoding this Header (see Format).
     Format Format
 }
 func (h *Header) FileInfo() os.FileInfo
 ------------------------------------------------------------------------------------
 **Description:
 A Header represents a single header in a tar archive. Some fields may not be populated.
 For forward compatibility, users that retrieve a Header from Reader.Next, mutate it
 in some ways, and then pass it back to Writer.WriteHeader should do so by creating
 a new Header and copying the fields that they are interested in preserving.
 FileInfo returns an os.FileInfo for the Header.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Header代表tar档案文件里的单个头, Header类型的某些字段可能未使用，定义说明如下：
 type Header struct {
    Name       string    // 记录头域的文件名
    Mode       int64     // 权限和模式位
    Uid        int       // 所有者的用户ID
    Gid        int       // 所有者的组ID
    Size       int64     // 字节数（长度）
    ModTime    time.Time // 修改时间
    Typeflag   byte      // 记录头的类型
    Linkname   string    // 链接的目标名
    Uname      string    // 所有者的用户名
    Gname      string    // 所有者的组名
    Devmajor   int64     // 字符设备或块设备的major number
    Devminor   int64     // 字符设备或块设备的minor number
    AccessTime time.Time // 访问时间
    ChangeTime time.Time // 状态改变时间
    Xattrs     map[string]string
 }
 1. FileInfo方法返回该Header对应的文件信息。（os.FileInfo类型）
 2. 可以使用部分参数声明和初始化来定义Header结构体。
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
