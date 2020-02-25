/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Int31n
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int31n(n int32) int32
 ------------------------------------------------------------------------------------
 **Description:
 Int31n returns, as an int32, a non-negative pseudo-random number in [0,n) from the
 default Source. It panics if n <= 0.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Int31n函数返回一个[0,n)之间的非负伪随机整数；
 2. 如果n<=0会panic。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Int31n()=", rand.Int31n(9999))
	fmt.Println("Int31n()=", rand.Int31n(9999))
}
