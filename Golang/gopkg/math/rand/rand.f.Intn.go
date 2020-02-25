/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Intn
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Intn(n int) int
 ------------------------------------------------------------------------------------
 **Description:
 Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the
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
	fmt.Println("Intn()=", rand.Intn(99999999))
	fmt.Println("Intn()=", rand.Intn(99999999))
}
