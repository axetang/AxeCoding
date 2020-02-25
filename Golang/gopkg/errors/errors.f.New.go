/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: errors
 **Element: errors.New()
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(text string) error
 ------------------------------------------------------------------------------------
 **Description:
 New returns an error that formats as the given text.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 查看errors包的源代码，New函数定义如下
 func New(text string) error {
	return &errorString{text}
 }
 即New()返回使用text字符串构建的private结构体errorString的指针*errors.errorString，定义如下：
 type errorString struct {
	s string
 }
 2. 该结构体有一个public的Error方法，返回该结构体的private字符串字段s
 func (e *errorString) Error() string {
	return e.s
 }
 即在使用New()构建出*errorString结构体，使用结构体的Error()方法获取结构体的private成员字符串s的值。
 3. errorString是errors包的private结构体，该结构体的方法Error()是public的，对于非导出成员，
 不再做进一步介绍，如需了解，可以查看gopkg的源码，呵呵。
 4. error是内置接口类型，在builtin包中定义，如下
 type error interface {
	Error() string
}
*************************************************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is a error created by errors.New()")
	if err != nil {
		fmt.Println(err)
	}
	err = errors.New("This is the 2nd error created by errors.New()")
	fmt.Println(err.Error())
	fmt.Printf("err's type is %T\n", err)

	//fmt.Errorf
	const name, id = "bimmler", 17
	err = fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
