/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.NewFloat
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewFloat(name string) *Float
 ------------------------------------------------------------------------------------
 **Description:
 1.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Float代表一个64位浮点数变量，满足Var接口;
 1. NewFloat函数创建一个*Float
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
)

func main() {
	fFloat := expvar.NewFloat("fFloat")
	fFloat.Set(3.14)
	fmt.Println(fFloat.Value())
}
