/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.IntsAreSorted
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IntsAreSorted(a []int) bool
 ------------------------------------------------------------------------------------
 **Description:
 IntsAreSorted tests whether a slice of ints is sorted in increasing order.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 函数IntsAreSorted判断是否参数a []int已经按照递增序排列。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var is = make([]int, 10)
	for i := 0; i < 10; i++ {
		is[i] = rand.Intn(100)
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	fmt.Println("IntsAreSorted?", sort.IntsAreSorted(is))
	sort.Ints(is)
	fmt.Println("-----Ints()-------")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	fmt.Println("IntsAreSorted?", sort.IntsAreSorted(is))

	s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sort.IntsAreSorted(s))
}
