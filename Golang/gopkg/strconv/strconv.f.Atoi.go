/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.Atoi
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Atoi(s string) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Atoi相当于ParseInt(s,10,0);

*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
}
