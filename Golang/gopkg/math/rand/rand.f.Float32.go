/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Float32
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Float32() float32
 ------------------------------------------------------------------------------------
 **Description:
 Float32 returns, as a float32, a pseudo-random number in [0.0,1.0) from the default
 Source.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("rand.Float32():", rand.Float32())
	fmt.Println("rand.Float32():", rand.Float32())
}
