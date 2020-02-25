/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Scanner
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Scanner interface {
    Scan(state ScanState, verb rune) error
 }
 ------------------------------------------------------------------------------------
 **Description:
 Scanner is implemented by any value that has a Scan method, which scans the input
 for the representation of a value and stores the result in the receiver, which
 must be a pointer to be useful. The Scan method is called for any argument to
 Scan, Scanf, or Scanln that implements it.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Scanner接口封装了Scan方法；
 2. Scan方法将输入扫描成值的表示，并将其结果存储到接收者中， 该接收者必须为可用的指针；
 3. Scan方法会被Scan、Scanf或Scanln的任何实现了它的实参所调用;
 4. ScanState接口定义参见fmt.i.ScanState.go
*************************************************************************************/
package main

func main() {
}
