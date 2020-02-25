/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Float
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Float struct {
     // contains filtered or unexported fields
	f uint64
 }
 func NewFloat(name string) *Float
 func (v *Float) Add(delta float64)
 func (v *Float) Set(value float64)
 func (v *Float) String() string
 func (v *Float) Value() float64
 ------------------------------------------------------------------------------------
 **Description:
 Float is a 64-bit float variable that satisfies the Var interface.
 Add adds delta to v.
 Set sets v to value.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Float代表一个64位浮点数变量，满足Var接口, Float只有一个非导出成员i int64;
 2. String方法满足了Var接口实现要求，返回Int的字符串字面量表达；
 3. Add方法实现将加到delta的v *Int的功能；
 4. Set方法给v *Int设置值value int64；
 5. Value方法返回v *Int的值。
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
	fFloat.Add(1.11)

	fmt.Println(fFloat.Value(), fFloat.String())
}
