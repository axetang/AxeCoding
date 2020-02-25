/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Unix
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Unix(sec int64, nsec int64) Time
 ------------------------------------------------------------------------------------
 **Description:
 Unix returns the local Time corresponding to the given Unix time, sec seconds and
 nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the
 range [0, 999999999]. Not all sec values have a corresponding time value. One such
 value is 1<<63-1 (the largest int64 value).
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Unix函数参数给定一个Unix时间：距离January 1， 1970 UTC的sec seconds和nsec nanoseconds；
 2. Unix函数返回这个相对时间的本地Local时间信息结构体Time；
 3.
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(10000000000, 100000000)
	fmt.Println(t)
}
