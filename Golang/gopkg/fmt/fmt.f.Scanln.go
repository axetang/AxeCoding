/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Scanln
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Scanln(a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Scanln is similar to Scan, but stops scanning at a newline and after the final item
 there must be a newline or EOF.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Scanln类似于Scan，但它在换行符处停止扫描，且最后的条目之后必须为换行符或EOF;
 2. 查看标准包源码，Fscan定义和实现如下
 func Scanln(a ...interface{}) (n int, err error) {
	return Fscanln(os.Stdin, a...)
 }
 3. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果。
*************************************************************************************/
package main

import (
	"fmt"
)

func main() {
	var i int
	var str string
	var f float64
	var b bool

	fmt.Scanln(&i, &str, &f, &b)
	fmt.Println(i, str, f, b)
}
