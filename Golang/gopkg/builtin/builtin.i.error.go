/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.error
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type error interface {
    Error() string
 }
 ------------------------------------------------------------------------------------
 **Description:
 The error built-in interface type is the conventional interface for representing an
 error condition, with the nil value representing no error.
 ------------------------------------------------------------------------------------
 **要点总结:
 error内建接口类型是表示错误情况的约定接口，nil值即表示没有错误。该接口唯一方法Error要求返回错误文本。
*************************************************************************************/
package main

import "fmt"

type MyError struct {
	ErrorText string
}

func (me *MyError) Error() string {
	return me.ErrorText
}

func main() {
	MyErr := new(MyError)
	MyErr.ErrorText = "my error!"
	fmt.Println(MyErr.Error())

}
