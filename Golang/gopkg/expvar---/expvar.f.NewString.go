/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.NewString
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewString(name string) *String
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewString方法创建一个*String变量
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
)

func main() {
	sStr := expvar.NewString("sStr")
	sStr.Set("Axe Tang")
	fmt.Println(sStr.Value())

}
