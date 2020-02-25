/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Uint64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Uint64() uint64
 ------------------------------------------------------------------------------------
 **Description:
 Uint64 returns a pseudo-random 64-bit value as a uint64 from the default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Uint64使用缺省源生成并返回一个64-bits的伪随机无符号整数。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Uint64()=", rand.Uint64())
	fmt.Println("Uint64()=", rand.Uint64())
}
