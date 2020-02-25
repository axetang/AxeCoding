/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.IsSorted
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func IsSorted(data Interface) bool
 ------------------------------------------------------------------------------------
 **Description:
 IsSorted reports whether data is sorted.
 ------------------------------------------------------------------------------------
 **要点总结:
 IsSorted判断参数data Interface已经排序。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var is = make(sort.IntSlice, 10, 20)
	for i := 0; i < 9; i++ {
		is[i] = rand.Intn(100)
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	is[9] = 50
	fmt.Printf("is[9]=%d\n", is[9])
	fmt.Println("IsSorted?", sort.IsSorted(is))
	fmt.Println("is.Len()=", is.Len())
	fmt.Println("is.Search(50)=", is.Search(50))
	fmt.Println("is.Search(65)=", is.Search(65))

	is.Sort()
	fmt.Println("IsSorted?", sort.IsSorted(is))
	fmt.Println("----Sort()----")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}

	fmt.Println("is.Search(50)=", is.Search(50))
	fmt.Println("is.Search(65)=", is.Search(65))

	fmt.Println("----Swap(8,9)----")
	is.Swap(8, 9)
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}

}
