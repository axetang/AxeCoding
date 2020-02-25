/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Errorf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Errorf(format string, a ...interface{}) error
 ------------------------------------------------------------------------------------
 **Description:
 Errorf formats according to a format specifier and returns the string as a value
 that satisfies error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看标准包源码，Errorf定义和实现如下
 func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
 }
 2. Errorf是通过errors.New+Sprintf函数来实现的，生成一个格式化字符串作为错误文本的error类型。
 *************************************************************************************/
package main

import (
	"fmt"
)

func main() {
	var str = "this is an error by axe!"
	err := fmt.Errorf("Hi, %s", str)
	fmt.Println(err)
	fmt.Println(err.Error())
}
