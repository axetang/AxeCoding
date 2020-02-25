/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.String
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // String is a string variable, and satisfies the Var interface.
 type String struct {
     // contains filtered or unexported fields
	s atomic.Value // string
 }
 func NewString(name string) *String
 func (v *String) Set(value string)
 func (v *String) String() string
 func (v *String) Value() string
 ------------------------------------------------------------------------------------
 **Description:
 String is a string variable, and satisfies the Var interface.
 String implements the Var interface. To get the unquoted string use Value.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. String结构体代表一个字符串；它实现了String方法，从而实现了Var接口；
 1. Set方法给v设定字符串值；
 2. Value()方法返回字符串的值；
 3. String()方法实现了Var接口，返回带双引号的字符串，JSON。
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
	fmt.Println(sStr.String())

}
