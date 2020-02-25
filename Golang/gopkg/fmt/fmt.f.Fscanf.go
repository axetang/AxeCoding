/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fscanf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Fscanf scans text read from r, storing successive space-separated values into
 successive arguments as determined by the format. It returns the number of items
 successfully parsed. Newlines in the input must match newlines in the format.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fscan定义和实现如下，该函数将从r中按照指定格式format读取值并依次赋予参数a切片成员；
 func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
	s, old := newScanState(r, false, false)
	n, err = s.doScanf(format, a)
	s.free(old)
	return
 }
 2. Fscanf调用了fmt包的私有函数newScanState和doScanf，具体可查阅标准包源代码；
 3. Fscanf在连续调用时，fmt的末尾要增加一个\n，否则第一次调用后后续的调用都不会执行，因为系统仍然以为
 上一次的输入读写扫描还没有结束，最终会读出异常值;
 4. 注意：在调用scan类函数时，因为scan函数的变参a切片是interface类型，所以实例化后的形参必须也是引用
 形式的变量，如&i，&str等，否则会执行出非常奇怪的结果。
 *************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	var i, j int
	var str string
	var f float64
	var b bool

	fmt.Fscanf(os.Stdin, "%d %s %f %t\n", &i, &str, &f, &b)
	fmt.Println(i, str, f, b)
	fmt.Fscanf(os.Stdin, "%d,%s,%f,%t\n", &j, &str, &f, &b)
	fmt.Println(j, str, f, b)

}
