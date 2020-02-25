/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.SearchInts
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SearchInts(a []int, x int) int
 ------------------------------------------------------------------------------------
 **Description:
 SearchInts searches for x in a sorted slice of ints and returns the index as
 specified by Search. The return value is the index to insert x if x is not
 present (it could be len(a)). The slice must be sorted in ascending order.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SearchInts在递增顺序的a中搜索x，返回x的索引；
 2. 如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
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

	sort.Ints(is)
	for i := 0; i < 10; i++ {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	fmt.Println("SearchInts(50)=", sort.SearchInts(is, 50))
	fmt.Println("SearchInts(68)=", sort.SearchInts(is, 68))
}
