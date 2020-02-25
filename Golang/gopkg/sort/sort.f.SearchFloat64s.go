/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.SearchFloat64s
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SearchFloat64s(a []float64, x float64) int
 ------------------------------------------------------------------------------------
 **Description:
 SearchFloat64s searches for x in a sorted slice of float64s and returns the index
 as specified by Search. The return value is the index to insert x if x is not
 present (it could be len(a)). The slice must be sorted in ascending order.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var fs = make([]float64, 10, 20)
	for i := 0; i < 9; i++ {
		fs[i] = rand.Float64() * 10
		fmt.Printf("fs[%d]=%f\n", i, fs[i])
	}
	fs[9] = 5.1
	fmt.Printf("is[9]=%f\n", fs[9])

	sort.Float64s(fs)
	for i, fsi := range fs {
		fmt.Println("fs[", i, "]=", fsi)
	}
	fmt.Println("SearchFloat64s(fs,5.0)=", sort.SearchFloat64s(fs, 5.1))
	fmt.Println("SearchFloat64s(fs,6.0)=", sort.SearchFloat64s(fs, 8.0))

}
