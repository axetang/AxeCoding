/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.WriterAt
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 WriterAt is the interface that wraps the basic WriteAt method.
 WriteAt writes len(p) bytes from p to the underlying data stream at offset off. It
 returns the number of bytes written from p (0 <= n <= len(p)) and any error
 encountered that caused the write to stop early. WriteAt must return a non-nil
 error if it returns n < len(p).
 If WriteAt is writing to a destination with a seek offset, WriteAt should not
 affect nor be affected by the underlying seek offset.
 Clients of WriteAt can execute parallel WriteAt calls on the same destination if
 the ranges do not overlap.
 Implementations must not retain p.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. WriterAt接口包装了基本的WriteAt方法;
 2. WriteAt从p中将len(p)个字节写入到偏移量为off的Writer数据流中。它返回从p中被写入的字节数n
 （0<=n<=len(p)）以及任何遇到的引起写入提前停止的错误。若WriteAt返回的n<len(p)，它就必须
 返回一个非nil的错误;
 3. 若WriteAt在有Seek设定偏移量的情况下对目标Writer进行写操作，则WriteAt应当既不影响Seek的
 offset偏移量，也不被这个偏移量所影响;
*************************************************************************************/
package main

func main() {
}
