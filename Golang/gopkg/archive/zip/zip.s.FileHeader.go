/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.FileHeader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type FileHeader struct {
     // Name is the name of the file.
     //
     // It must be a relative path, not start with a drive letter (such as "C:"),
     // and must use forward slashes instead of back slashes. A trailing slash
     // indicates that this file is a directory and should have no data.
     //
     // When reading zip files, the Name field is populated from
     // the zip file directly and is not validated for correctness.
     // It is the caller's responsibility to sanitize it as
     // appropriate, including canonicalizing slash directions,
     // validating that paths are relative, and preventing path
     // traversal through filenames ("../../../").
     Name string

     // Comment is any arbitrary user-defined string shorter than 64KiB.
     Comment string

     // NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
     //
     // By specification, the only other encoding permitted should be CP-437,
     // but historically many ZIP readers interpret Name and Comment as whatever
     // the system's local character encoding happens to be.
     //
     // This flag should only be set if the user intends to encode a non-portable
     // ZIP file for a specific localized region. Otherwise, the Writer
     // automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
     NonUTF8 bool

     CreatorVersion uint16
     ReaderVersion  uint16
     Flags          uint16

     // Method is the compression method. If zero, Store is used.
     Method uint16

     // Modified is the modified time of the file.
     //
     // When reading, an extended timestamp is preferred over the legacy MS-DOS
     // date field, and the offset between the times is used as the timezone.
     // If only the MS-DOS date is present, the timezone is assumed to be UTC.
     //
     // When writing, an extended timestamp (which is timezone-agnostic) is
     // always emitted. The legacy MS-DOS date field is encoded according to the
     // location of the Modified time.
     Modified     time.Time
     ModifiedTime uint16 // Deprecated: Legacy MS-DOS date; use Modified instead.
     ModifiedDate uint16 // Deprecated: Legacy MS-DOS time; use Modified instead.

     CRC32              uint32
     CompressedSize     uint32 // Deprecated: Use CompressedSize64 instead.
     UncompressedSize   uint32 // Deprecated: Use UncompressedSize64 instead.
     CompressedSize64   uint64
     UncompressedSize64 uint64
     Extra              []byte
     ExternalAttrs      uint32 // Meaning depends on CreatorVersion
 }
 func (h *FileHeader) FileInfo() os.FileInfo
 func (h *FileHeader) ModTime() time.Time
 func (h *FileHeader) Mode() (mode os.FileMode)
 func (h *FileHeader) SetModTime(t time.Time)
 func (h *FileHeader) SetMode(mode os.FileMode)
 ------------------------------------------------------------------------------------
 **Description:
 FileHeader describes a file within a zip file. See the zip spec for details.
 FileInfo returns an os.FileInfo for the FileHeader.
 ModTime returns the modification time in UTC using the legacy ModifiedDate and
 ModifiedTime fields.
 Deprecated: Use Modified instead.
 Mode returns the permission and mode bits for the FileHeader.
 SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields to the given
 time in UTC.
 Deprecated: Use Modified instead.
 SetMode changes the permission and mode bits for the FileHeader.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. FileHeader描述zip文件中的一个文件。参见zip的定义获取细节:
 type FileHeader struct {
    // Name是文件名，它必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔符
    Name string
    CreatorVersion     uint16
    ReaderVersion      uint16
    Flags              uint16
    Method             uint16
    ModifiedTime       uint16 // MS-DOS时间
    ModifiedDate       uint16 // MS-DOS日期
    CRC32              uint32
    CompressedSize     uint32 // 已弃用；请使用CompressedSize64
    UncompressedSize   uint32 // 已弃用；请使用UncompressedSize64
    CompressedSize64   uint64
    UncompressedSize64 uint64
    Extra              []byte
    ExternalAttrs      uint32 // 其含义依赖于CreatorVersion
    Comment            string
 }
 1. FileInfo返回一个根据h的信息生成的os.FileInfo。
 2. Mode返回h的权限和模式位。
 3. SetMode修改h的权限和模式位。
 4. ModTime返回最近一次修改的UTC时间。（精度2s）
 5. SetModTime将ModifiedTime和ModifiedDate字段设置为给定的UTC时间（精度2s）.
*************************************************************************************/
package main

func main() {
}
