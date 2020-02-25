/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Error
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Error interface {
     error
     Timeout() bool   // Is the error a timeout?
     Temporary() bool // Is the error temporary?
 }
 ------------------------------------------------------------------------------------
 **Description:
 An Error represents a network error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Error接口表示一个网络错误信息；
 2. Error定义说明如下：
 type Error interface {
    error
    Timeout() bool   // 错误是否为超时？
    Temporary() bool // 错误是否是临时的？
 }
*************************************************************************************/
package main

func main() {
}
