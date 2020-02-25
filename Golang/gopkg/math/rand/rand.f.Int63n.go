/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Int63n
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int63n(n int64) int64
 func Intn(n int) int
 ------------------------------------------------------------------------------------
 **Description:
 Int63n returns, as an int64, a non-negative pseudo-random number in [0,n) from the
 default Source. It panics if n <= 0.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Int63n(9999999999)=", rand.Int63n(9999999999))
	fmt.Println("Int63n(9999999999)=", rand.Int63n(9999999999))
}
