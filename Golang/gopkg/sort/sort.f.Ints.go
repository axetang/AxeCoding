/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Ints
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Ints(a []int)
 ------------------------------------------------------------------------------------
 **Description:
 Ints sorts a slice of ints in increasing order.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Ints函数给参数int切片排递增序。
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
	sort.Ints(is)
	fmt.Println("-----Ints()-------")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
}
