/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: RuneScanner
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type RuneScanner interface {
    RuneReader
    UnreadRune() error
 }
 ------------------------------------------------------------------------------------
 **Description:
 1. RuneScanner is the interface that adds the UnreadRune method to the basic ReadRune
 method.
 2. UnreadRune causes the next call to ReadRune to return the same rune as the
 previous call to ReadRune. It may be an error to call UnreadRune twice without an
 intervening call to ReadRune.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneScanner接口将UnreadRune方法添加到基本的ReadRune方法。
 2. UnreadRune使下一次调用ReadRune返回的符文与上一次调用ReadRune返回的相同。 调用UnreadRune两次
 而中间没有调用ReadRune的话就会返回错误。
*************************************************************************************/
package main

func main() {
}
