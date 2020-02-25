/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Scan
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Scan(a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Scan scans text read from standard input, storing successive space-separated values
 into successive arguments. Newlines count as space. It returns the number of items
 successfully scanned. If that is less than the number of arguments, err will report
 why.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Scan扫描从标准输入中读取的文本，并将连续由空格分隔的值存储为连续的实参,换行符计为空格,它返回
 成功扫描的条目数,若它少于实参数，err就会报告原因;
 2. 查看标准包源码，Scan定义和实现如下
 func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
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

	fmt.Scan(&i, &f, &str, &b)
	//fmt.Scan(&str)
	fmt.Println(i, str, f, b)
}
