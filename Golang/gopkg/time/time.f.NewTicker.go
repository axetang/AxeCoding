/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.NewTicker
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewTicker(d Duration) *Ticker
 ------------------------------------------------------------------------------------
 **Description:
 NewTicker returns a new Ticker containing a channel that will send the time with a
 period specified by the duration argument. It adjusts the intervals or drops ticks
 to make up for slow receivers. The duration d must be greater than zero; if not,
 NewTicker will panic. Stop the ticker to release associated resources.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewTicker函数返回一个新的Ticker结构体，它包含一个channel，这个channel会按照参数d Duration指定
 的时间间隔发送时间信息；
 2. 它可以调整时间间隔或取消tick发送从而构成更慢的接受者。参数d Duration必须大于
 0，否则，NewTicker函数会panic，停止ticker并释放相关资源；
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
