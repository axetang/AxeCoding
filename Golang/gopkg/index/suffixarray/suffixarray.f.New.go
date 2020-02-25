/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: suffixarray
 **Element: suffixarray.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(data []byte) *Index
 ------------------------------------------------------------------------------------
 **Description:
 New creates a new Index for data. Index creation time is O(N*log(N)) for N = len(data).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. New函数使用给订的参数data []byte数据生成一个*Index结构体，时间复杂度O(N*log(N))。
*************************************************************************************/
package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	data := []byte("aaaa")
	index := suffixarray.New(data)
	fmt.Println(index)
}
