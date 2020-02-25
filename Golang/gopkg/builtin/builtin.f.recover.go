/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.recover
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func recover() interface{}
 ------------------------------------------------------------------------------------
 **Description:
 The recover built-in function allows a program to manage behavior of a panicking
 goroutine. Executing a call to recover inside a deferred function (but not any
 function called by it) stops the panicking sequence by restoring normal execution
 and retrieves the error value passed to the call of panic. If recover is called
 outside the deferred function it will not stop a panicking sequence. In this case,
 or when the goroutine is not panicking, or if the argument supplied to panic was
 nil, recover returns nil. Thus the return value from recover reports whether the
 goroutine is panicking.
 ------------------------------------------------------------------------------------
 **要点总结:
 recover内建函数使程序可以管理恐慌中的 goroutine 。recover在推迟函数（其它函数总是返回nil）
 中调用，通过执行recover停止恐慌过程序列并取回传至panic的参数值，恢复正常的执行。若recover不在
 推迟函数中被调用，它将不会停止恐慌过程序列。在此情况下或该goroutine不在恐慌过程中，或提供给panic
 的实参为nil ，recover就会返回nil。因此recover的返回值就报告了该goroutine是否进入恐慌过程序列。
*************************************************************************************/
package main

import (
	"fmt"
)

//Try ddd
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}

//say ddd
func say(s string) {
	fmt.Println(s)
}

func main() {
	say("Hello")
	Try(
		func() {
			panic("World")
		},
		func(e interface{}) {
			fmt.Println("catch", e)
			fmt.Printf("e type : %T\n", e)
		})
	say("end")
}
