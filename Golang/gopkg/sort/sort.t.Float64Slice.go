/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Float64Slice
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Float64Slice []float64
 func (p Float64Slice) Len() int
 func (p Float64Slice) Less(i, j int) bool
 func (p Float64Slice) Search(x float64) int
 func (p Float64Slice) Sort()
 func (p Float64Slice) Swap(i, j int)
 ------------------------------------------------------------------------------------
 **Description:
 Float64Slice attaches the methods of Interface to []float64, sorting in increasing order (not-a-number values are treated as less than other values).
 Search returns the result of applying SearchFloat64s to the receiver and x.
 Sort is a convenience method.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Float64Slice给[]int添加方法以满足Interface接口，以便排序为递增序列；
 2. Len、Less、Swap方法实现了排序需要的功能；
 3. Search方法对方法接受者p执行sort.SearchInts()函数，参数是x float64。如果x在p中存在，则返回x
 对应的元素的index；如果找不到x，则返回x应该插入的位置；
 4. 从以上Search方法执行要求来看，要执行search方法，必须保证SliceInt实例已经Sort，否则执行结果错误。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var fs = make(sort.Float64Slice, 10, 20)
	for i := 0; i < 9; i++ {
		fs[i] = rand.Float64() * 10
		fmt.Printf("is[%d]=%f\n", i, fs[i])
	}
	fs[9] = 5.0
	fmt.Printf("is[9]=%f\n", fs[9])

	fmt.Println("fs.Len()=", fs.Len())
	fmt.Println("fs.Search(5.0)=", fs.Search(5.0))
	fmt.Println("fs.Search(6.8)=", fs.Search(6.8))

	fs.Sort()
	fmt.Println("----Sort()----")
	for i := range fs {
		fmt.Printf("fs[%d]=%f\n", i, fs[i])
	}

	fmt.Println("fs.Search(5.0)=", fs.Search(5.0))
	fmt.Println("fs.Search(6.8)=", fs.Search(6.8))

	fmt.Println("----Swap(8,9)----")
	fs.Swap(8, 9)
	for i := range fs {
		fmt.Printf("fs[%d]=%f\n", i, fs[i])
	}

}
