/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Search
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Search(n int, f func(int) bool) int
 ------------------------------------------------------------------------------------
 **Description:
 Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true, assuming that on the range [0, n), f(i) == true implies f(i+1) == true. That is, Search requires that f is false for some (possibly empty) prefix of the input range [0, n) and then true for the (possibly empty) remainder; Search returns the first true index. If there is no such index, Search returns n. (Note that the "not found" return value is not -1 as in, for instance, strings.Index.) Search calls f(i) only for i in the range [0, n).
 A common use of Search is to find the index i for a value x in a sorted, indexable data structure such as an array or slice. In this case, the argument f, typically a closure, captures the value to be searched for, and how the data structure is indexed and ordered.
 For instance, given a slice data sorted in ascending order, the call Search(len(data), func(i int) bool { return data[i] >= 23 }) returns the smallest index i such that data[i] >= 23. If the caller wants to find whether 23 is in the slice, it must test data[i] == 23 separately.
 Searching data sorted in descending order would use the <= operator instead of the >= operator.
 To complete the example above, the following code tries to find the value x in an integer slice data sorted in ascending order:
 x := 23
 i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
 if i < len(data) && data[i] == x {
 	// x is present at data[i]
 } else {
 	// x is not present in data,
 	// but i is the index where it would be inserted.
 }
 As a more whimsical example, this program guesses your number:
 func GuessingGame() {
 	var s string
 	fmt.Printf("Pick an integer from 0 to 100.\n")
 	answer := sort.Search(100, func(i int) bool {
 		fmt.Printf("Is your number <= %d? ", i)
 		fmt.Scanf("%s", &s)
 		return s != "" && s[0] == 'y'
 	})
 	fmt.Printf("Your number is %d.\n", answer)
 }
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。也就是说，Search函数希望
 f在输入位于区间[0, n)的前面某部分（可以为空）时返回假，而在输入位于剩余至结尾的部分（可以为空）时
 返回真；Search函数会返回满足f(i)==true的最小值i。如果没有该值，函数会返回n。注意，未找到时的返回
 值不是-1，这一点和strings.Index等函数不同。Search函数只会用区间[0, n)内的值调用f;
 2. 一般使用Search找到值x在插入一个有序的、可索引的数据结构时，应插入的位置。这种情况下，参数f（通常
 是闭包）会捕捉应搜索的值和被查询的数据集;
 3. 例如，给定一个递增顺序的切片，调用
 Search(len(data), func(i int) bool { return data[i] >= 23 })
 会返回data中最小的索引i满足data[i] >= 23。如果调用者想要知道23是否在切片里，它必须另外检查
 data[i] == 23;
 4. 搜索递减顺序的数据时，应使用<=运算符代替>=运算符。
*************************************************************************************/
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	println("-----GuessingGame()---------")
	GuessingGame()
}
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
