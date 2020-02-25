/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Int31
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int31() int32
 ------------------------------------------------------------------------------------
 **Description:
 Int31 returns a non-negative pseudo-random 31-bit integer as an int32 from the
 default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Int31()=", rand.Int31())
	fmt.Println("Int31()=", rand.Int31())
}
