/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Buffers
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Buffers [][]byte
 func (v *Buffers) Read(p []byte) (n int, err error)
 func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 Buffers contains zero or more runs of bytes to write.
 On certain machines, for certain types of connections, this is optimized into an
 OS-specific batch write operation (such as "writev").
*************************************************************************************/
package main

func main() {
}
