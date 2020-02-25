/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.Error
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Error struct {
     Op  string
     URL string
     Err error
 }
 func (e *Error) Error() string
 func (e *Error) Temporary() bool
 func (e *Error) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 Error reports an error and the operation and URL that caused it.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Error结构体报告一个错误，以及导致该错误发生的URL和操作；
 1. Error方法返回错误信息字符串；
 2. Temporary方法判断错误是否是一个临时错误，具体功能待查看源码后更新；
 3. Timeout方法判断错误是否是一个超时错误。
*************************************************************************************/
package main

func main() {
}
