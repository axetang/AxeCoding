/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Seed
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Seed(seed int64)
 ------------------------------------------------------------------------------------
 **Description:
 Seed uses the provided seed value to initialize the default Source to a deterministic
 state. If Seed is not called, the generator behaves as if seeded by Seed(1). Seed
 values that have the same remainder when divided by 2^31-1 generate the same
 pseudo-random sequence. Seed, unlike the Rand.Seed method, is safe for concurrent
 use.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 以提供的参数作为种子值,来初始化随机产生器.如果Seed没有被调用,那么随机数调用前默认调用Seed(1).
 2. 小贴士,一般用传入time.Now().UnixNano()给Seed函数来实现不同运行次数看到不同结果的目的.
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10
	i := 0
	for i < n {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Int())
		i += 1
	}
}
