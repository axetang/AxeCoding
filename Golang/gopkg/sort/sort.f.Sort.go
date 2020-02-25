/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Sort
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sort(data Interface)
 ------------------------------------------------------------------------------------
 **Description:
 Sort sorts data. It makes one call to data.Len to determine n, and O(n*log(n))
 calls to data.Less and data.Swap. The sort is not guaranteed to be stable.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sort函数对参数data Interface进行排序；
 2. 参数data实现了Interface接口。
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

	fmt.Println("is.Len()=", is.Len())
	fmt.Println("is.Search(50)=", is.Search(50))
	fmt.Println("is.Search(65)=", is.Search(65))

	is.Sort()
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
