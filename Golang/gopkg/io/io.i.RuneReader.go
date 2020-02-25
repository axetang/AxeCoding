/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.RuneReader
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type RuneReader interface {
    ReadRune() (r rune, size int, err error)
 }
 ------------------------------------------------------------------------------------
 **Description:
 1. RuneReader is the interface that wraps the ReadRune method.
 2. ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its
 size in bytes. If no character is available, err will be set.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. RuneReader 接口包装了 ReadRune 方法;
 2. ReadRune读取单个用UTF-8编码的Unicode字符，并返回该符文及其字节大小。若没有字符可用，就会置为err。
*************************************************************************************/
package main

func main() {
}
