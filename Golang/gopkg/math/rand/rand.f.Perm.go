/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Perm
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Perm(n int) []int
 ------------------------------------------------------------------------------------
 **Description:
 Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
 [0,n) from the default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Perm函数返回一个元素个数是n的int slice,里面的元素是0到n-1的整数无重复的随机排列.
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10
	i := 0
	for i < n {
		x := rand.Perm(5)
		fmt.Println(x)
		i += 1
	}
}
