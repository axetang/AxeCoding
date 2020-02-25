/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.ScanState
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type ScanState interface {
    // ReadRune reads the next rune (Unicode code point) from the input.
    // If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
    // return EOF after returning the first '\n' or when reading beyond
    // the specified width.
    ReadRune() (r rune, size int, err error)
    // UnreadRune causes the next call to ReadRune to return the same rune.
    UnreadRune() error
    // SkipSpace skips space in the input. Newlines are treated appropriately
    // for the operation being performed; see the package documentation
    // for more information.
    SkipSpace()
    // Token skips space in the input if skipSpace is true, then returns the
    // run of Unicode code points c satisfying f(c).  If f is nil,
    // !unicode.IsSpace(c) is used; that is, the token will hold non-space
    // characters. Newlines are treated appropriately for the operation being
    // performed; see the package documentation for more information.
    // The returned slice points to shared data that may be overwritten
    // by the next call to Token, a call to a Scan function using the ScanState
    // as input, or when the calling Scan method returns.
    Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
    // Width returns the value of the width option and whether it has been set.
    // The unit is Unicode code points.
    Width() (wid int, ok bool)
    // Because ReadRune is implemented by the interface, Read should never be
    // called by the scanning routines and a valid implementation of
    // ScanState may choose always to return an error from Read.
    Read(buf []byte) (n int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 ScanState represents the scanner state passed to custom scanners. Scanners may
 do rune-at-a-time scanning or ask the ScanState to discover the next
 space-delimited token.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ReadRune从输入中读取下一个符文（Unicode码点）。若它在 Scanln、Fscanln 或Sscanln中被
 调用，ReadRune() 就会在返回第一个“\n”或读取至指定宽度后返回 EOF;
 2. SkipSpace跳过输入中的空格。换行符会被视作空格，除非该扫描操作为Scanln、Fscanln或Sscanln，
 在这种情况下，换行符会被视作EOF;
 3. Token会在skipSpace为true时从输入中跳过空格，然后返回一系列满足f(c)的Unicode码点c;
 若f为nil，就会使用!unicode.IsSpace(c)；即该标记会保留非空格字符。换行符会被视作空格，
 除非该扫描操作为Scanln、Fscanln 或 Sscanln，否则在这种情况下，换行符会被视作EOF;
 4. 所返回的指向共享数据的切片可能会在下次调用 Token 时被覆盖，Scan 函数的调用会将ScanState
 用作输入，或当调用Scan方法返回时也会;
 5. Width返回宽度选项的值，并判断它是否被设置。其单元为Unicode码点;
 6. 由于ReadRune被此接口实现，因此Read不应当被扫描功能调用，而ScanState的有效实现可选择总是
 从Read中返回错误。
*************************************************************************************/
package main

func main() {
}
