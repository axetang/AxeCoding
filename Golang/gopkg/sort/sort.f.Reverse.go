/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Reverse
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Reverse(data Interface) Interface
 ------------------------------------------------------------------------------------
 **Description:
 Reverse returns the reverse order for data.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Reverse函数返回data Interface的反序；
 2. 从执行结果看，Reverse必须使用Sort来再排一次序才能得到正确结果，即
	 sort.Sort(sort.Reverse(data))
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

	is.Sort()
	fmt.Println("----Sort()----")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}

	sort.Sort(sort.Reverse(is))
	fmt.Println("----Reverse()----")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}

	fmt.Println("\n-----Test Reverse--------")
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
