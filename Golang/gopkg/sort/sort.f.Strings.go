/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Strings
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Strings(a []string)
 ------------------------------------------------------------------------------------
 **Description:
 Strings sorts a slice of strings in increasing order.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var letter = make([]byte, 26, 26)
	for i := 0; i < 26; i++ {
		letter[i] = byte('a' + i)
		//fmt.Println(string(letter[i]))
	}

	var ss = make([]string, 10, 20)
	for i := 0; i < 9; i++ {
		ss[i] = string(letter[rand.Intn(26)]) + string(letter[rand.Intn(25-i)])
		fmt.Printf("is[%d]=%s\n", i, ss[i])
	}
	ss[9] = "mn"
	fmt.Printf("ss[9]=%s\n", ss[9])

	sort.Strings(ss)
	fmt.Println("----Sort()----")
	for i := range ss {
		fmt.Printf("ss[%d]=%s\n", i, ss[i])
	}

}
