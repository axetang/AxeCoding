/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: list
 **Element: list.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New() *List
 New returns an initialized list.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. New函数创建一个空链表，链表的长度为0，开头和末尾节点都是`nil`；
 2.

*************************************************************************************/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e := l.Front()        // 取出链表开头的节点，即节点a
	fmt.Println(e == nil) // 输出：true
}
