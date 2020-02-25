/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    SeekStart   = 0 // seek relative to the origin of the file
    SeekCurrent = 1 // seek relative to the current offset
    SeekEnd     = 2 // seek relative to the end
 )
 var EOF = errors.New("EOF")
 var ErrClosedPipe = errors.New("io: read/write on closed pipe")
 var ErrNoProgress = errors.New("multiple Read calls return no data or error")
 var ErrShortBuffer = errors.New("short buffer")
 var ErrShortWrite = errors.New("short write")
 var ErrUnexpectedEOF = errors.New("unexpected EOF")
 ------------------------------------------------------------------------------------
 **Description:
 1. 3 consts for Seek whence values;
 2. EOF is the error returned by Read when no more input is available. Functions
 should return EOF only to signal a graceful end of input. If the EOF occurs
 unexpectedly in a structured data stream, the appropriate error is either
 ErrUnexpectedEOF or some other error giving more detail;
 3. ErrClosedPipe is the error used for read or write operations on a closed pipe;
 4. ErrNoProgress is returned by some clients of an io.Reader when many calls to
 Read have failed to return any data or error, usually the sign of a broken
 io.Reader implementation;
 5. ErrShortBuffer means that a read required a longer buffer than was provided.
 6. ErrShortWrite means that a write accepted fewer bytes than requested but failed
 to return an explicit error;
 7. ErrUnexpectedEOF means that EOF was encountered in the middle of reading a
 fixed-size block or data structure.
 ------------------------------------------------------------------------------------
 **要点总结:
1. EOF是Read在没有更多输入可用时返回的错误。函数应当只为标志出优雅的输入结束而返回EOF。若EOF在结构
化数据流中出现意外，适当的错误不是ErrUnexpectedEOF，就是一些其它能给出更多详情的错误;
2. ErrClosedPipe错误用于在已关闭的管道上进行读取或写入操作;
3. io.Reader的使用者返回ErroNoProgress错误，表示很多读操作调用失败且不能返回任何数据或错误，通常
出现在这个io.Reader的实现崩溃时；
4. ErrShortBuffer意为所需读取的缓存比提供的长;
5. ErrShortWrite意为接受写入的字节比请求的少，但无法返回一个明确的错误;
6. ErrUnexpectedEOF意为在读取固定大小的块或数据结构过程中遇到EOF.
*************************************************************************************/
package main

func main() {
}
