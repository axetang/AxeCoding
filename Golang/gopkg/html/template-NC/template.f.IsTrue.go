/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.IsTrue
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsTrue(val interface{}) (truth, ok bool)
 ------------------------------------------------------------------------------------
 **Description:
 IsTrue reports whether the value is 'true', in the sense of not the zero of its
 type, and whether the value has a meaningful truth value. This is the definition
 of truth used by if and other such actions.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 此函数执行结果有点怪，需要了解源码后更新。
*************************************************************************************/
package main

import (
	"fmt"
	"html/template"
)

func main() {
	t, _ := template.IsTrue("True")
	fmt.Println("True is true?", t)
	t, _ = template.IsTrue("false")
	fmt.Println("false is true?", t)
	t, _ = template.IsTrue("False")
	fmt.Println("False is true?", t)
	t, _ = template.IsTrue(5)
	fmt.Println("5 is true?", t)
	t, _ = template.IsTrue(0)
	fmt.Println("0 is true?", t)
	t, _ = template.IsTrue(-1)
	fmt.Println("-1 is true?", t)
	t, _ = template.IsTrue(nil)
	fmt.Println("nil is true?", t)
}
