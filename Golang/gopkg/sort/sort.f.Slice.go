/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Slice
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Slice(slice interface{}, less func(i, j int) bool)
 ------------------------------------------------------------------------------------
 **Description:
 Slice sorts the provided slice given the provided less function.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Slice函数对slice参数进行排序，依据提供的less func(i,j int)函数；
 2. 通常less是个闭包
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var is = make([]int, 10, 20)
	for i := 0; i < 9; i++ {
		is[i] = rand.Intn(100)
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	is[9] = 50
	fmt.Printf("is[9]=%d\n", is[9])

	fmt.Println("SearchInts(50)=", sort.SearchInts(is, 50))
	fmt.Println("SearchInts(68)=", sort.SearchInts(is, 68))

	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
	for i := 0; i < 10; i++ {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	fmt.Println("SearchInts(50)=", sort.SearchInts(is, 50))
	fmt.Println("SearchInts(68)=", sort.SearchInts(is, 68))
}
