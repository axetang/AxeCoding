/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Sleep
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Sleep(d Duration)
 ------------------------------------------------------------------------------------
 **Description:
 Sleep pauses the current goroutine for at least the duration d. A negative or zero
 duration causes Sleep to return immediately.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Sleep函数停止当前goroutine一段时间，时间长度为d Duration；
 2. 负值或零值会导致Sleep立即结束并返回
*************************************************************************************/
package main

import "time"

func main() {
	time.Sleep(100 * time.Millisecond)
	time.Sleep(-100 * time.Second)
}
