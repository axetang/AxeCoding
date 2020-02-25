/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Weekday
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Weekday int
 const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
 )
 func (d Weekday) String() string
 ------------------------------------------------------------------------------------
 **Description:
 A Weekday specifies a day of the week (Sunday = 0, ...).
 String returns the English name of the day ("Sunday", "Monday", ...).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Weekday类型定义一个星期中的每一天，通过一个常数系列来指定，每个常数的类型都是Weekday；
 2. String方法返回Weekday常数的字符串字面量。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Sunday.String())
	fmt.Println(time.Sunday)
	fmt.Printf("%d %d %d\n", time.Monday, time.Wednesday, time.Sunday)
}
