/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.GoStringer
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type GoStringer interface {
    GoString() string
 }
 ------------------------------------------------------------------------------------
 **Description:
 GoStringer is implemented by any value that has a GoString method, which defines
 the Go syntax for that value. The GoString method is used to print values passed
 as an operand to a %#v format.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. GoStringer接口由任何拥有GoString方法的值所实现，该方法定义了该值的Go语法格式;
 2. GoString方法用于打印作为操作数传至%#v进行格式化的值。
*************************************************************************************/
package main

func main() {
}
