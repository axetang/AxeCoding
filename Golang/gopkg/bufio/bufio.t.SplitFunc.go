/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bufio
**Element: bufio.SplitFunc
**Type: type
------------------------------------------------------------------------------------
**Definition:
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
------------------------------------------------------------------------------------
**Description:
SplitFunc is the signature of the split function used to tokenize the input. The
arguments are an initial substring of the remaining unprocessed data and a flag,
atEOF, that reports whether the Reader has no more data to give. The return
values are the number of bytes to advance the input and the next token to return
to the user, if any, plus an error, if any.
Scanning stops if the function returns an error, in which case some of the input
may be discarded.
Otherwise, the Scanner advances the input. If the token is not nil, the Scanner
returns it to the user. If the token is nil, the Scanner reads more data and
continues scanning; if there is no more data--if atEOF was true--the Scanner
returns. If the data does not yet hold a complete token, for instance if it has
no newline while scanning lines, a SplitFunc can return (0, nil, nil) to signal
the Scanner to read more data into the slice and try again with a longer slice
starting at the same point in the input.
The function is never called with an empty data slice unless atEOF is true. If
at EOF is true, however, data may be non-empty and, as always, holds unprocessed
text.
------------------------------------------------------------------------------------
**要点总结:
1. SplitFunc类型代表用于对输出作词法分析的分割函数;
2. 参数data是尚未处理的数据的一个开始部分的切片， 参数atEOF表示是否Reader接口不能提供更多的
数据。 返回值是解析位置前进的字节数，将要返回给调用者的token切片， 以及可能遇到的错误。如果数据
不足以（保证）生成一个完整的token， 例如需要一整行数据但data里没有换行符， SplitFunc可以返回
(0, nil, nil)来告诉Scanner读取更多的数据 写入切片然后用从同一位置起始、长度更长的切片再试一
次（调用SplitFunc类型函数）;
3. 如果返回值err非nil，扫描将终止并将该错误返回给Scanner的调用者;
4. 除非atEOF为真，永远不会使用空切片data调用SplitFunc类型函数。 然而，如果atEOF为真，data
却可能是非空的、且包含着未处理的文本。
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a new error!")
	if err != nil {
		fmt.Print(err)
	}
}
