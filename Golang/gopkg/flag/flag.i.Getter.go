/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Getter
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Getter interface {
    Value
    Get() interface{}
 }
 ------------------------------------------------------------------------------------
 **Description:
 Getter is an interface that allows the contents of a Value to be retrieved.
 It wraps the Value interface, rather than being part of it, because it
 appeared after Go 1 and its compatibility rules. All Value types provided
 by this package satisfy the Getter interface.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Getter接口包含两个成员，一个是接口Value，一个是Get方法；
 2. Get方法使得一个Value接口实例的值可以被获取；
 3. flag包中的所有Value类型实例都实现了Getter接口；
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

type MyValue struct {
	i int
}

func (mv *MyValue) String() string {
	s := strconv.Itoa(mv.i)
	return s
}

func (mv *MyValue) Set(s string) error {
	mv.i, _ = strconv.Atoi(s)
	return nil
}

func (mv *MyValue) Get() interface{} {
	return mv.i
}

func main() {
	var mv MyValue
	mv.Set("5")
	fmt.Println("Get", mv.Get())
	fmt.Println("String", mv.String())
}
