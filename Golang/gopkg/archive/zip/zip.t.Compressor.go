/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.Compressor
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Compressor func(w io.Writer) (io.WriteCloser, error)
 ------------------------------------------------------------------------------------
 **Description:
 A Compressor returns a new compressing writer, writing to w. The WriteCloser's Close method must be used to flush pending data to w. The Compressor itself must be safe to invoke from multiple goroutines simultaneously, but each returned writer will be used only by one goroutine at a time.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Compressor函数类型会返回一个io.WriteCloser，该接口会将数据压缩后写入提供的接口;
 1. 关闭时，应将缓冲中的数据刷新到下层接口中。
*************************************************************************************/
package main

func main() {
}
