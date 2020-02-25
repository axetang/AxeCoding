/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Stringer
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Stringer interface {
    String() string
 }
 ------------------------------------------------------------------------------------
 **Description:
 Stringer is implemented by any value that has a String method, which defines the
 “native” format for that value. The String method is used to print values passed
 as an operand to any format that accepts a string or to an unformatted printer
 such as Print.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Stringer接口由任何拥有String方法的值所实现，该方法定义了该值的“原生”格式。 String方法
 用于打印值，该值可作为操作数传至任何接受字符串的格式，或像Print这样的未格式化打印器;
 2. 任何实现了String方法的类型，在通过Print系列函数打印时会被自动调用。
*************************************************************************************/
package main

func main() {
}
