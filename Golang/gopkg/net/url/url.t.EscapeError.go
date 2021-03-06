/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.EscapeError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type EscapeError string
 func (e EscapeError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 0. EscapeError类型记录转码错误信息；
 1. EscapeError实现了error接口的Error方法，返回错误信息。
*************************************************************************************/
package main

func main() {
}
