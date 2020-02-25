/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fscanln
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Fscanln is similar to Fscan, but stops scanning at a newline and after the final
 item there must be a newline or EOF.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fscanln定义和实现如下
 func Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
	s, old := newScanState(r, false, true)
	n, err = s.doScan(a)
	s.free(old)
	return
 }
 2. Fscanln调用了fmt包的私有函数newScanState和doScan，具体可查阅标准包源代码；
 3. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	str := "abc"
	f := 3.14
	t := true

	fmt.Fscanln(os.Stdin, &i, &str, &f, &t)
	println(i, str, f, t)

	fmt.Fscanln(os.Stdin, &i, &str, &f, &t)
	println(i, str, f, t)
}
