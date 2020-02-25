/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.Decompressor
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Decompressor func(r io.Reader) io.ReadCloser
 ------------------------------------------------------------------------------------
 **Description:
 A Decompressor returns a new decompressing reader, reading from r. The ReadCloser's
 Close method must be used to release associated resources. The Decompressor itself
 must be safe to invoke from multiple goroutines simultaneously, but each returned
 reader will be used only by one goroutine at a time.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Decompressor函数类型会返回一个io.ReadCloser， 该接口的Read方法会将读取自提供的接口的数据
 提前解压缩;
 1. 程序员有责任在读取结束时关闭该io.ReadCloser。
*************************************************************************************/
package main

func main() {
}
