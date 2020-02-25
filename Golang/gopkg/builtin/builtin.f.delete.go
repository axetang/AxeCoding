/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.deletes
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func delete(m map[Type]Type1, key Type)
 ------------------------------------------------------------------------------------
 **Description:
 The delete built-in function deletes the element with the specified key (m[key])
 from the map. If m is nil or there is no such element, delete is a no-op.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. delete内建函数按照指定的键将元素从映射中删除。如果元素不存在，delete为空操作。如果m为nil，
 delete引发panic错误。
*************************************************************************************/
package main

import "fmt"

//Vertex ddd
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var nilmap map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	delete(m, "key")
	fmt.Println(m)
	delete(m, "Bell Labs")
	fmt.Println(m)

	delete(nilmap, "key")

}
