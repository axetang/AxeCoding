/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Scanf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Scanf(format string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Scanf扫描从标准输入中读取的文本，并将连续由空格分隔的值存储为连续的实参，其格式由format决定,它
 返回成功扫描的条目数, 若返回的条目数小于实参数，则会报告错误原因err;
 2. 查看标准包源码，Scanf定义和实现如下
 func Scanf(format string, a ...interface{}) (n int, err error) {
	return Fscanf(os.Stdin, format, a...)
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

	fmt.Scanf("%d %s %f %t", &i, &str, &f, &b)
	fmt.Println(i, str, f, b)
}
