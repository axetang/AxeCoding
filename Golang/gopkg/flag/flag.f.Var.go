/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Var
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Var(value Value, name string, usage string)
------------------------------------------------------------------------------------
 **Description:
 Var defines a flag with the specified name and usage string. The type and
 value of the flag are represented by the first argument, of type Value,
 which typically holds a user-defined implementation of Value. For instance,
 the caller could create a flag that turns a comma-separated string into a
 slice of strings by giving the slice the methods of Value; in particular,
 Set would decompose the comma-separated string into the slice.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. Var函数构建一个自定义Value接口类型的flag；
 2. name和usage参数指定了flag的名字和usage。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

type MyValue struct {
	s string
}

func (mv *MyValue) String() string {
	return mv.s
}
func (mv *MyValue) Set(s string) error {
	mv.s = s
	return nil
}

//这里如果编程mv=&MyValue{}，下面的flag.Var函数调用，就要去掉mv前面的引用符号&
//反之亦反
var mv = MyValue{}

func init() {
	flag.Var(&mv, "name", "enter your name.")
}
func main() {
	flag.Parse()
	mv.Set("Axe Tang")
	fmt.Println(mv.String())
}
