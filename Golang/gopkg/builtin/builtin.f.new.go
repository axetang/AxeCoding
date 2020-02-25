/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: builtin
**Element: builtin.new
**Type: func
------------------------------------------------------------------------------------
**Definition:
func new(Type) *Type
------------------------------------------------------------------------------------
**Description:
The new built-in function allocates memory. The first argument is a type, not a
value, and the value returned is a pointer to a newly allocated zero value of that
type.
------------------------------------------------------------------------------------
**要点总结:
1. new内建函数用于给值类型分配内存，返回指向类型的指针；
2. new内建函数第一个实参为类型，而非值，其返回值为指向该类型新分配的零值的指针；
3. 值类型包括基本类型，struct，array和自定义类型；
*************************************************************************************/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	var d Vertex
	v := new(Vertex)
	d.X = 8
	fmt.Println(v, d)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
