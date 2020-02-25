/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ring
 **Element: ring.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(n int) *Ring
 ------------------------------------------------------------------------------------
 **Description:
 New creates a ring of n elements.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. New函数创建一个具有n个元素的环形链表。
*************************************************************************************/
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(10)
	fmt.Println(r.Len() == 10) // 输出：true
}
