/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.IntSlice
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type IntSlice []int
 func (p IntSlice) Len() int
 func (p IntSlice) Less(i, j int) bool
 func (p IntSlice) Search(x int) int
 func (p IntSlice) Sort()
 func (p IntSlice) Swap(i, j int)
 ------------------------------------------------------------------------------------
 **Description:
 IntSlice attaches the methods of Interface to []int, sorting in increasing order.
 Search returns the result of applying SearchInts to the receiver and x.
 Sort is a convenience method.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. IntSlice给[]int添加方法以满足Interface接口，以便排序为递增序列；
 2. Len、Less、Swap方法实现了排序需要的功能；
 3. Search方法对方法接受者p执行sort.SearchInts()函数，参数是x int。如果x在p中存在，则返回x
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
