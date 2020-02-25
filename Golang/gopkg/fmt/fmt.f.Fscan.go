/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fscan
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fscan(r io.Reader, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Fscan scans text read from r, storing successive space-separated values into
 successive arguments. Newlines count as space. It returns the number of items
 successfully scanned. If that is less than the number of arguments, err will report
 why.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fscan定义和实现如下
 func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	s, old := newScanState(r, true, false)
	n, err = s.doScan(a)
	s.free(old)
	return
 }
 2. Fscan调用了fmt包的私有函数newScanState和doScan，具体可查阅标准包源代码；
 3. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	var i int
	var str string
	var f float64
	var b bool

	fmt.Fscan(os.Stdin, &i, &str, &f, &b)
	fmt.Println(i, str, f, b)

}
