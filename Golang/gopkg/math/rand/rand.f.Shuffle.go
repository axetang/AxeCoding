/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Shuffle
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Shuffle(n int, swap func(i, j int))
 ------------------------------------------------------------------------------------
 **Description:
 Shuffle pseudo-randomizes the order of elements using the default Source. n is the
 number of elements. Shuffle panics if n < 0. swap swaps the elements with indexes
 i and j.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Shuffle函数使用缺省源来伪随机化元素顺序；
 2. 参数n是元素数量；
 3. swap函数用i和j索引交换元素；
 4. 此函数功能还需要深入了解后刷新。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println("n =", len(words))
	fmt.Println(words)
}
