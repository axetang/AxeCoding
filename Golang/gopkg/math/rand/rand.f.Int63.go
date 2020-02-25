/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: rand
**Element: rand.Int63
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Int63() int64
------------------------------------------------------------------------------------
**Description:
Int63 returns a non-negative pseudo-random 63-bit integer as an int64 from the
default Source.
------------------------------------------------------------------------------------
**要点总结:
1. Int63返回一个63位的非负伪随机数，从缺省Source；
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Int63()=", rand.Int63())
	fmt.Println("Int63()=", rand.Int63())
}
