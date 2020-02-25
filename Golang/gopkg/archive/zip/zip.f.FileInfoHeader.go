/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.FileInfoHeader
 **Type: func
 ------------------------------------------------------------------------------------
 func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
 ------------------------------------------------------------------------------------
 **Description:
 FileInfoHeader creates a partially-populated FileHeader from an os.FileInfo. Because
 os.FileInfo's Name method returns only the base name of the file it describes, it
 may be necessary to modify the Name field of the returned header to provide the
 full path name of the file. If compression is desired, callers should set the
 FileHeader.Method field; it is unset by default.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FileInfoHeader返回一个根据fi填写了部分字段的Header;
 2. 因为os.FileInfo接口的Name方法只返回它描述的文件的无路径名，有可能需要将返回值的Name字段
 修改为文件的完整路径名;
*************************************************************************************/
package main

func main() {
}
