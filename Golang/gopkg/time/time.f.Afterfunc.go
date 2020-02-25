/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.AfterFunc
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 func AfterFunc(d Duration, f func()) *Timer
 ------------------------------------------------------------------------------------
 **Description:
 AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
 It returns a Timer that can be used to cancel the call using its Stop method.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. AfterFunc函数经过参数时间d Duration之后在自己的goroutine中调用参数函数f；
 2. 可以调用AfterFunc函数返回的Timer的Stop方法来停止Timer。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	f := func() {
		fmt.Println("Time out")
	}
	t := time.AfterFunc(5*time.Second, f)

	t.Stop()
}
