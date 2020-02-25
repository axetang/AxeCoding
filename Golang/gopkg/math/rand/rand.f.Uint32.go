/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Uint32
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Uint32() uint32
 ------------------------------------------------------------------------------------
 **Description:
 Uint32 returns a pseudo-random 32-bit value as a uint32 from the default Source.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Uint32使用缺省源生成并返回一个32-bits的伪随机无符号整数。
**************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Uint32()=", rand.Uint32())
	fmt.Println("Uint32()=", rand.Uint32())
}
