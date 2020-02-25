/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Sscan
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sscan(str string, a ...interface{}) (n int, err error)
 func Sscanf(str string, format string, a ...interface{}) (n int, err error)
 func Sscanln(str string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Sscan scans the argument string, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.
 Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.
 Sscanln is similar to Sscan, but stops scanning at a newline and after the final item there must be a newline or EOF.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sscan扫描实参string，并将连续由空格分隔的值存储为连续的实参, 换行符计为空格, 它返回成功扫描的
 条目数。若它少于实参数，err 就会报告原因;
 2. 查看标准包源码，Sscan定义和实现如下
 func Sscan(str string, a ...interface{}) (n int, err error) {
	return Fscan((*stringReader)(&str), a...)
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

	str = "1 3.14 true"
	fmt.Sscan(str, &i, &f, &b)
	fmt.Println(i, f, b)
}
