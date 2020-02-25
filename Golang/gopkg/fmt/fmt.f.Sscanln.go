/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Sscanln
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sscanln(str string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Sscanln is similar to Sscan, but stops scanning at a newline and after the final item there must be a newline or EOF.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sscanln类似于Sscan，但它在换行符处停止扫描，且最后的条目之后必须为换行符或EOF. 如下程序试验，
 上述规则不起作用，应该说Sscanln读完了要读的数据就自动停止了，而不会等待换行符或EOF，或者提前遇到
 了换行符会即时终止扫描;
 2. 查看标准包源码，Scanln定义和实现如下
 func Sscanln(str string, a ...interface{}) (n int, err error) {
	return Fscanln((*stringReader)(&str), a...)
 }
 3. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果。
*************************************************************************************/
package main

import (
	"fmt"
)

func main() {
	var i1, i2, i3, i4 int
	var str string
	var f1, f2, f3, f4 float64
	var b1, b2, b3, b4 bool

	str = "1 3.14 true"
	fmt.Sscanln(str, &i1, &f1, &b1)
	fmt.Println(i1, f1, b1)

	str = "1 3.14 true\n"
	fmt.Sscanln(str, &i2, &f2, &b2)
	fmt.Println(i2, f2, b2)

	str = "1 3.14 true 888 "
	fmt.Sscanln(str, &i3, &f3, &b3)
	fmt.Println(i3, f3, b3)

	str = "1 \n3.14 \n true 888 "
	fmt.Sscanln(str, &i4, &f4, &b4)
	fmt.Println(i4, f4, b4)
}
