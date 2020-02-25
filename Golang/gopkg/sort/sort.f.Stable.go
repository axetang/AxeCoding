/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Stable
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Stable(data Interface)
 ------------------------------------------------------------------------------------
 **Description:
 Stable sorts data while keeping the original order of equal elements.
 It makes one call to data.Len to determine n, O(n*log(n)) calls to data.Less and
 O(n*log(n)*log(n)) calls to data.Swap.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Stable函数对参数data Interface进行排序，并确保相等元素的原始顺序；
 2. data是Interface实例，因此实现了Len,Less,Swap方法。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var is = make(sort.IntSlice, 10)
	for i := 0; i < 10; i++ {
		if i == 5 || i == 9 {
			is[5] = 50
			is[9] = 50
		} else {
			is[i] = rand.Intn(100)
		}
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
	sort.Stable(is)
	fmt.Println("-----Stable-------")
	for i := range is {
		fmt.Printf("is[%d]=%d\n", i, is[i])
	}
}
