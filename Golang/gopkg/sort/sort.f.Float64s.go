/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Float64s(a []float64)
 func Float64sAreSorted(a []float64) bool
 func Ints(a []int)
 func IntsAreSorted(a []int) bool
 func IsSorted(data Interface) bool
 func Search(n int, f func(int) bool) int
 ------------------------------------------------------------------------------------
 **Description:
 Float64s sorts a slice of float64s in increasing order (not-a-number values are treated as less than other values).
 Float64sAreSorted tests whether a slice of float64s is sorted in increasing order (not-a-number values are treated as less than other values).
 Ints sorts a slice of ints in increasing order.
 IntsAreSorted tests whether a slice of ints is sorted in increasing order.
 IsSorted reports whether data is sorted.
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
*************************************************************************************/
package main

func main() {
}
