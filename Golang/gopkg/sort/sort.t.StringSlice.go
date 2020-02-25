/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.StringSlice
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type StringSlice []string
 func (p StringSlice) Len() int
 func (p StringSlice) Less(i, j int) bool
 func (p StringSlice) Search(x string) int
 func (p StringSlice) Sort()
 func (p StringSlice) Swap(i, j int)
 ------------------------------------------------------------------------------------
 **Description:
 StringSlice attaches the methods of Interface to []string, sorting in increasing order.
 Search returns the result of applying SearchStrings to the receiver and x.
 Sort is a convenience method.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. StringSlice给[]int添加方法以满足Interface接口，以便排序为递增序列；
 2. Len、Less、Swap方法实现了排序需要的功能；
 3. Search方法对方法接受者p执行sort.SearchInts()函数，参数是x string。如果x在p中存在，则返回x
 对应的元素的index；如果找不到x，则返回x应该插入的位置；
 4. 从以上Search方法执行要求来看，要执行search方法，必须保证StringSlice实例已经Sort，否则执行
 结果错误。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var letter = make([]byte, 26, 26)
	for i := 0; i < 26; i++ {
		letter[i] = byte('a' + i)
		//fmt.Println(string(letter[i]))
	}

	var ss = make(sort.StringSlice, 10, 20)
	for i := 0; i < 9; i++ {
		ss[i] = string(letter[rand.Intn(26)]) + string(letter[rand.Intn(25-i)])
		fmt.Printf("is[%d]=%s\n", i, ss[i])
	}
	ss[9] = "mn"
	fmt.Printf("ss[9]=%s\n", ss[9])

	fmt.Println("ss.Len()=", ss.Len())
	fmt.Println("ss.Search(\"mn\")=", ss.Search("mn"))
	fmt.Println("ss.Search(\"op\")=", ss.Search("op"))

	ss.Sort()
	fmt.Println("----Sort()----")
	for i := range ss {
		fmt.Printf("ss[%d]=%s\n", i, ss[i])
	}

	fmt.Println("ss.Search(\"mn\")=", ss.Search("mn"))
	fmt.Println("ss.Search(\"op\")=", ss.Search("op"))

	fmt.Println("----Swap(8,9)----")
	ss.Swap(8, 9)
	for i := range ss {
		fmt.Printf("ss[%d]=%s\n", i, ss[i])
	}

}
