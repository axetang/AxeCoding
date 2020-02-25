/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Int
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Int() int
 ------------------------------------------------------------------------------------
 **Description:
 Int returns a non-negative pseudo-random int from the default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
 Int从缺省Source返回一个非负伪随机整数；
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("rand.Int():", rand.Int())
	fmt.Println("rand.Int():", rand.Int())
}
