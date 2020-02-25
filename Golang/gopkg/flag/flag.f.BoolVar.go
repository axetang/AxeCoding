/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.BoolVar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func BoolVar(p *bool, name string, value bool, usage string)
------------------------------------------------------------------------------------
 **Description:
 BoolVar defines a bool flag with specified name, default value, and usage
 string. The argument p points to a bool variable in which to store the
 value of the flag.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. BoolVar定义一个bool flag命令行参数，通过参数指定了参数的name string，value bool和usage string。
 2. 第一个参数p是一个指向bool变量的指针，它保存了这个bool flag的值。
*************************************************************************************/
package main

import "flag"

var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "vector", true, "is vector?")
}

func main() {
	flag.Parse()

	if h && v {
		flag.Usage()
	}
}
