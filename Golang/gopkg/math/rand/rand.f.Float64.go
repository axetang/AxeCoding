/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Float64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Float64() float64
 ------------------------------------------------------------------------------------
 **Description:
 Float64 returns, as a float64, a pseudo-random number in [0.0,1.0) from the default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("rand.Float64():", rand.Float64())
	fmt.Println("rand.Float64():", rand.Float64())
}
