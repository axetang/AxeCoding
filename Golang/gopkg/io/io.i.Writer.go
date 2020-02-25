/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.Writer
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Writer interface {
    Write(p []byte) (n int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 Writer is the interface that wraps the basic Write method.
 Write writes len(p) bytes from p to the underlying data stream. It returns the
 number of bytes written from p (0 <= n <= len(p)) and any error encountered that
 caused the write to stop early. Write must return a non-nil error if it returns
 n < len(p). Write must not modify the slice data, even temporarily.
 Implementations must not retain p.
 ------------------------------------------------------------------------------------
 **要点总结:1.
*************************************************************************************/
package main

func main() {
}
