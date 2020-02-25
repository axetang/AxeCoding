/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: expvar
**Element: expvar.NewMap
**Type: func
------------------------------------------------------------------------------------
**Definition:
func NewMap(name string) *Map
------------------------------------------------------------------------------------
**Description:
------------------------------------------------------------------------------------
**要点总结:
1. NewMap函数创建并返回一个*Map，其导出的名字叫做name string。
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
)

func main() {
	m := expvar.NewMap("mymap")
	fmt.Println(m)

}
