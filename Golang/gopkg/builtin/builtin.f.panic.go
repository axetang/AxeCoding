/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.panic
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func panic(v interface{})
 ------------------------------------------------------------------------------------
 **Description:
 The panic built-in function stops normal execution of the current goroutine. When a
 function F calls panic, normal execution of F stops immediately. Any functions whose
 execution was deferred by F are run in the usual way, and then F returns to its
 caller. To the caller G, the invocation of F then behaves like a call to panic,
 terminating G's execution and running any deferred functions. This continues until
 all functions in the executing goroutine have stopped, in reverse order. At that
 point, the program is terminated and the error condition is reported, including the
 value of the argument to panic. This termination sequence is called panicking and
 can be controlled by the built-in function recover.
 ------------------------------------------------------------------------------------
 **要点总结:
 panic内建函数停止当前goroutine的正常执行。当函数F调用panic，函数F的执行被终止，并且F中的延迟函数
 会正常执行，然后F返回到调用它的地方。对于调用者G，F的行为就像调用了panic，即终止G的执行并运行所有
 被推迟的函数。这一过程会到该程序中所有函数都按相反的顺序停止执行之后。此时，该程序会被终止，并产生
 错误报告，包括引发该panic的实参值。此终止序列称为恐慌过程，可以通过内建函数recover控制。
*************************************************************************************/
package main

import "fmt"

func f() {
	defer func() {
		fmt.Println("f() defer")
	}()
	fmt.Println("f() before panic")
	panic(0)
	// 此后的代码不会被执行
	fmt.Println("after panic")
	defer func() {
		fmt.Println("f() defer agin")
	}()
}

func g() {
	defer func() {
		fmt.Println("g() defer")
	}()
	f()
	// 此后的代码不会被执行
	fmt.Println("after call f()")
}

func main() {
	g()
}
