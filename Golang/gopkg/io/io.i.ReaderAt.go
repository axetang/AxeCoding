/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.ReaderAt
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 ReaderAt is the interface that wraps the basic ReadAt method. ReadAt reads len(p)
 bytes into p starting at offset off in the underlying input source. It returns the
 number of bytes read (0 <= n <= len(p)) and any error encountered.
 When ReadAt returns n < len(p), it returns a non-nil error explaining why more bytes
 were not returned. In this respect, ReadAt is stricter than Read.
 Even if ReadAt returns n < len(p), it may use all of p as scratch space during the
 call. If some data is available but not len(p) bytes, ReadAt blocks until either all
 the data is available or an error occurs. In this respect ReadAt is different from
 Read.
 If the n = len(p) bytes returned by ReadAt are at the end of the input source, ReadAt
 may return either err == EOF or err == nil.
 If ReadAt is reading from an input source with a seek offset, ReadAt should not
 affect nor be affected by the underlying seek offset.
 Clients of ReadAt can execute parallel ReadAt calls on the same input source.
 Implementations must not retain p.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ReaderAt接口包装了基本的ReadAt方法;
 2. ReadAt方法从输入源的偏移量off处开始，将len(p)个字节读取到p中。它返回读取的字节数n(0<=n<=len(p)）
 以及任何遇到的错误；
 3. 当ReadAt返回的n<len(p)时，它就会返回一个非nil的错误来解释为什么没有返回更多的字节。在这一点上，
 ReadAt比Read更严格。即使ReadAt返回的n<len(p)，它也会在调用过程中使用p的全部作为暂存空间。 若一
 些数据可用但不到len(p)字节，ReadAt就会阻塞直到所有数据都可用或产生一个错误。在这一点上ReadAt
 不同于Read;
 4. 若n=len(p)个字节在输入源的的结尾处由ReadAt返回，ReadAt不是返回err==EOF就是返回err==nil；
 5. 若ReadAt在设定Seek偏移量offset情况下从输入源读取，ReadAt应当既不影响Seek的offset偏移量
 也不被它所影响;
 6. 实现必须不保留p（这个的确切含义要进一步了解下再刷新说明）。
*************************************************************************************/
package main

func main() {
}
