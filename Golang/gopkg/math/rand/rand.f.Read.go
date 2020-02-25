/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Read
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Read(p []byte) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Read generates len(p) random bytes from the default Source and writes them into p.
 It always returns len(p) and a nil error. Read, unlike the Rand.Read method, is
 safe for concurrent use.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Read从缺省Source生成len(p)个字节的随机数并写入p；
 2. 并行安全。
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	p := []byte{1, 2, 3, 4, 5}
	fmt.Println(p)
	n, _ := rand.Read(p)
	fmt.Println(p, n)

}
