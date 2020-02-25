/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    // MaxScanTokenSize is the maximum size used to buffer a token
    // unless the user provides an explicit buffer with Scan.Buffer.
    // The actual maximum token size may be smaller as the buffer
    // may need to include, for instance, a newline.
    MaxScanTokenSize = 64 * 1024
 )
 var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
 )
 var (
    ErrTooLong         = errors.New("bufio.Scanner: token too long")
    ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative
    advance count")
    ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance
    count beyond input")
 )
 var ErrFinalToken = errors.New("final token")
 ------------------------------------------------------------------------------------
 **Description:
 ErrFinalToken is a special sentinel error value. It is intended to be returned by
 a Split function to indicate that the token being delivered with the error is the
 last token and scanning should stop after this one. After ErrFinalToken is
 received by Scan, scanning stops with no error. The value is useful to stop
 processing early or when it is necessary to deliver a final empty token. One
 could achieve the same behavior with a custom error value but providing one here
 is tidier. See the emptyFinalToken example for a use of this value.
 ------------------------------------------------------------------------------------
 **要点总结:

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
