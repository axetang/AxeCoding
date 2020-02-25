/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.LimitedReader
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type LimitedReader struct {
    R   Reader // underlying reader
    N   int64  // max bytes remaining
 }
 func (l *LimitedReader) Read(p []byte) (n int, err error)
 Description:
 A LimitedReader reads from R but limits the amount of data returned to just N bytes.
 Each call to Read updates N to reflect the new amount remaining. Read returns EOF
 when N <= 0 or when the underlying R returns EOF.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 结构体LimitedReader通过实现Read方法实现了Reader接口，任何实现Reader接口的变量都可以成为结构体成员R；
 2. 每次调用该结构体的Read方法会更新结构体成员N的值，用以反映R还剩余多少个字节尚未读取；
 3. 当N<=0或者R返回EOF的时候，Read方法返回EOF；
 4. 结构体LimitedReader的Read方法实现时，调用了结构体成员R的Read方法，这也是结构体包含R的必然。
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
