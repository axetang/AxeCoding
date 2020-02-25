/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.State
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type State interface {
    // Write is the function to call to emit formatted output to be printed.
    Write(b []byte) (n int, err error)
    // Width returns the value of the width option and whether it has been set.
    Width() (wid int, ok bool)
    // Precision returns the value of the precision option and whether it has been set.
    Precision() (prec int, ok bool)
    // Flag reports whether the flag c, a character, has been set.
    Flag(c int) bool
 }
 ------------------------------------------------------------------------------------
 **Description:
 State represents the printer state passed to custom formatters. It provides
 access to the io.Writer interface plus information about the flags and options
 for the operand's format specifier.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. State 表示传递给格式化器的打印器的状态。 它提供了访问 io.Writer 接口及关于标记的信息，
 2. 以及操作数的格式说明符选项;
 3. Write 函数用于打印出已格式化的输出;
 4. Width 返回宽度选项的值以及它是否已被设置;
 5. Precision 返回精度选项的值以及它是否已被设置;
 6. Flag 返回标记 c（一个字符）是否已被设置。
*************************************************************************************/
package main

func main() {
}
