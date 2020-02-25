/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package:rand
 **Element: rand.NewSource
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewSource(seed int64) Source
 ------------------------------------------------------------------------------------
 **Description:
 NewSource returns a new pseudo-random Source seeded with the given value. Unlike the
 default Source used by top-level functions, this source is not safe for concurrent
 use by multiple goroutines.
 ------------------------------------------------------------------------------------
 **要点总结:
 该函数主要返回一个指定种子的随机数产生器.
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

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < n {
		fmt.Println(r.Int())
		i += 1
	}

}
