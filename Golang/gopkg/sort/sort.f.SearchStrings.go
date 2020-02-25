/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.SearchStrings
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SearchStrings(a []string, x string) int
 ------------------------------------------------------------------------------------
 **Description:
 SearchStrings searches for x in a sorted slice of strings and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SearchStrings在递增顺序的a中搜索x，返回x的索引；
 2. 如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
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

	var ss = make([]string, 10, 20)
	for i := 0; i < 9; i++ {
		ss[i] = string(letter[rand.Intn(26)]) + string(letter[rand.Intn(25-i)])
		fmt.Printf("is[%d]=%s\n", i, ss[i])
	}
	ss[9] = "mn"
	fmt.Printf("ss[9]=%s\n", ss[9])

	fmt.Println("sort.SearchString(\"mn\")=", sort.SearchStrings(ss, "mn"))
	fmt.Println("sort.SearchString(\"op\")=", sort.SearchStrings(ss, "op"))

	sort.Strings(ss)
	fmt.Println("----Sort()----")
	for i := range ss {
		fmt.Printf("ss[%d]=%s\n", i, ss[i])
	}

	fmt.Println("sort.SearchString(\"mn\")=", sort.SearchStrings(ss, "mn"))
	fmt.Println("sort.SearchString(\"op\")=", sort.SearchStrings(ss, "op"))

}
