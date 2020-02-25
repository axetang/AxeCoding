/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.append
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func append(slice []Type, elems ...Type) []Type
 ------------------------------------------------------------------------------------
 **Description:
 The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:
	 slice = append(slice, elem1, elem2)
	 slice = append(slice, anotherSlice...)
 As a special case, it is legal to append a string to a byte slice, like this:
	 slice = append([]byte("hello "), "world"...)
 ------------------------------------------------------------------------------------
 **要点总结:
 1. append内建函数将元素追加到切片的末尾。若它有足够的容量，其目标就会重新切片以容纳新的元素。否则就会
 分配一个新的基本数组。append返回更新后的切片,因此必须存储追加后的结果，通常为包含该切片自身的变量：
 	slice = append(slice, elem1, elem2)
	slice = append(slice, anotherSlice...)
*************************************************************************************/
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:3]
	c := append(s, 6)
	fmt.Printf("&a=%p,&s=%p,&c=%p\n", &a, &s, &c)
	fmt.Printf("cap a=%d,cap s=%d,cap c=%d\n", cap(a), cap(s), cap(c))
	fmt.Printf("len a=%d,len s=%d,len c=%d\n", len(a), len(s), len(c))
	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("s ==", s[:cap(s)])
	fmt.Println("c ==", c[:cap(c)])

	d := append(a, 6)
	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("d ==", d[:cap(d)])
	fmt.Printf("cap a=%d,cap s=%d,cap c=%d cap d=%d\n", cap(a), cap(s), cap(c), cap(d))
	fmt.Printf("len a=%d,len s=%d,len c=%d,len d=%d\n", len(a), len(s), len(c), len(d))
}
