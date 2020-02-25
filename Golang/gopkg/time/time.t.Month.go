/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Month
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Month int
 const (
    January Month = 1 + iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
 )
 func (m Month) String() string
 ------------------------------------------------------------------------------------
 **Description:
 A Month specifies a month of the year (January = 1, ...).
 String returns the English name of the month ("January", "February", ...).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Month类型的元类型为int，在Month类型下定义了12个常量，表示一年的十二个月。
 2. String方法返回Month类型里常熟的字符串表达。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	y, m, d := time.Now().Date()
	fmt.Println(y, m, d)
}
