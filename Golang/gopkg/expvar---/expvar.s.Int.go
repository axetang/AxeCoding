/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Int
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 type Int struct {
    // contains filtered or unexported fields
	i int64
 }
 func (v *Int) Add(delta int64)
 func (v *Int) Set(value int64)
 func (v *Int) String() string
 func (v *Int) Value() int64
 ------------------------------------------------------------------------------------
 **Description:
 Int is a 64-bit integer variable that satisfies the Var interface.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Int代表一个64位整数变量，满足Var接口, Int只有一个非导出成员i int64;
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
	iInt := expvar.NewInt("iIntName")
	iInt.Set(100)
	fmt.Println(iInt.Value(), iInt.String())
	iInt.Add(200)
	fmt.Println(iInt.Value(), iInt.String())
}
