/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.SliceIsSorted
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
 ------------------------------------------------------------------------------------
 **Description:
 SliceIsSorted tests whether a slice is sorted.
 The function panics if the provided interface is not a slice.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SliceIsSorted函数判断参数slice interface{}是否已经按照函数参数less排序；
 2. less是一个闭包。
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
	fmt.Println("Slice is Sorted?", sort.SliceIsSorted(is, func(i, j int) bool { return is[i] < is[j] }))
	fmt.Println("Slice is Sorted?", sort.SliceIsSorted(is, func(i, j int) bool { return is[i] > is[j] }))
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
	fmt.Println("Slice is Sorted?", sort.SliceIsSorted(is, func(i, j int) bool { return is[i] < is[j] }))
	fmt.Println("Slice is Sorted?", sort.SliceIsSorted(is, func(i, j int) bool { return is[i] > is[j] }))
	for i := 0; i < 10; i++ {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	fmt.Println("SearchInts(50)=", sort.SearchInts(is, 50))
	fmt.Println("SearchInts(68)=", sort.SearchInts(is, 68))
}
