/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Fprintln
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
 ------------------------------------------------------------------------------------
 **Description:
 Fprintln formats using the default formats for its operands and writes to w. Spaces
 are always added between operands and a newline is appended. It returns the number
 of bytes written and any write error encountered.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Fprintln定义和实现如下,Fprintln把参数a切片中的值按照缺省格式写入io.Writer
 接口实例，操作符之间自动添加空格，最后增加一个换行，并返回写入的字节数
 func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
 }
 2. Fprintln调用了fmt包的私有函数newPrinter和doPrintln，查阅标准包源代码，两个函数定义如下。其中，
 doPrintln函数是pp类型的方法，pp类型定义也一并列出如下；
 func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.fmt.init(&p.buf)
	return p
}
func (p *pp) doPrintln(a []interface{}) {
	for argNum, arg := range a {
		if argNum > 0 {
			p.buf.WriteByte(' ')
		}
		p.printArg(arg, 'v')
	}
	p.buf.WriteByte('\n')
}
type pp struct {
	buf buffer

	// arg holds the current item, as an interface{}.
	arg interface{}

	// value is used instead of arg for reflect values.
	value reflect.Value

	// fmt is used to format basic items such as integers or strings.
	fmt fmt

	// reordered records whether the format string used argument reordering.
	reordered bool
	// goodArgNum records whether the most recent reordering directive was valid.
	goodArgNum bool
	// panicking is set by catchPanic to avoid infinite panic, recover, panic, ... recursion.
	panicking bool
	// erroring is set when printing an error string to guard against calling handleMethods.
	erroring bool
}
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	str := "abc"
	t := true

	fmt.Fprintln(os.Stdout, i, str, t)
}
