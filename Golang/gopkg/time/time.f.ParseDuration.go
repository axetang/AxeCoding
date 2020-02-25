/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.ParseDuration
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseDuration(s string) (Duration, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseDuration parses a duration string. A duration string is a possibly signed
 sequence of decimal numbers, each with optional fraction and a unit suffix, such
 as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms",
 "s", "m", "h".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseDuration方法分析duration字符串s string，并将其解析为一个Duration返回；
 2. duration字符串s中合法的时间单位是ns，us，ms，s，m，h。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	complex1, _ := time.ParseDuration("1h10m10s100ms100us1000ns")
	complex2, _ := time.ParseDuration("1h15m30.918273645s")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Println(complex1)
	fmt.Println(complex2)
	fmt.Printf("there are %.0f seconds in %v\n", complex.Seconds(), complex)
	fmt.Printf("there are %.0f seconds in %v\n", complex1.Seconds(), complex1)
	fmt.Printf("there are %.0f seconds in %v\n", complex2.Seconds(), complex2)
}
