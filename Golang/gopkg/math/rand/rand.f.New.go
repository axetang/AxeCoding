/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package:rand
 **Element: rand.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(src Source) *Rand
 ------------------------------------------------------------------------------------
 **Description:
 New returns a new Rand that uses random values from src to generate other random
 values.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. New函数返回一个新的Rand结构体实例,并以src作为随机值产生器;
 2.
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10
	i := 0

	r := rand.New(rand.NewSource(64))
	for i < n {
		fmt.Println(r.Int())
		i += 1
	}
}
