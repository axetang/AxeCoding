/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Sscanf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sscanf(str string, format string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sscanf扫描实参string，并将连续由空格分隔的值存储为连续的实参，其格式由format决定, 它返回成功
 解析的条目数;
 2. 查看标准包源码，Sscan定义和实现如下

 3. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果;
 4. 对于浮点数，不能提供精度值，否则会读取错误，如下程序注释部分；
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
	//fmt.Sscanf(str, "%d %5.2f %t", &i, &f, &b)
	fmt.Sscanf(str, "%d %f %t", &i, &f, &b)
	fmt.Println(i, f, b)
}
