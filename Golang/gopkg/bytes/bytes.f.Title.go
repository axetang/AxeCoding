/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: bytes
**Element: bytes.Title
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Title(s []byte) []byte
------------------------------------------------------------------------------------
**Description:
Title returns a copy of s with all Unicode letters that begin words mapped to their
title case.
------------------------------------------------------------------------------------
**要点总结:
1. Title返回s的一个副本，把s中每个单词的首字母改为Unicode字符的大写.
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("hello, world, i'm axe!")
	fmt.Println(string(bytes.Title(s)))
	fmt.Println(string(s))
}
