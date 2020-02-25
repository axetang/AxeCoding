/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Date
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
 ------------------------------------------------------------------------------------
 **Description:
 Date returns the Time corresponding to
 yyyy-mm-dd hh:mm:ss + nsec nanoseconds
 in the appropriate zone for that time in the given location.
 The month, day, hour, min, sec, and nsec values may be outside their usual ranges and
 will be normalized during the conversion. For example, October 32 converts to November
 1.
 A daylight savings time transition skips or repeats times. For example, in the United
 States, March 13, 2011 2:15am never occurred, while November 6, 2011 1:15am occurred
 twice. In such cases, the choice of time zone, and therefore the time, is not
 well-defined. Date returns a time that is correct in one of the two zones involved in
 the transition, but it does not guarantee which.
 Date panics if loc is nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Date函数返回按照"yyyy-mm-dd hh:mm:ss +nsec nanoseconds"格式的时间结构体Time；
 2. 注意：
	1）参数month类型为time.Month结构体，指定月份；
	2）参数loc类型*time.Location结构体指针，指定时区。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)
	fmt.Printf("Go launched at local: %s\n", t.Local())
}
