/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Ticker
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Ticker struct {
     C <-chan Time // The channel on which the ticks are delivered.
     // contains filtered or unexported fields
 }
 func (t *Ticker) Stop()
 ------------------------------------------------------------------------------------
 **Description:
 A Ticker holds a channel that delivers `ticks' of a clock at intervals.
 NewTicker returns a new Ticker containing a channel that will send the time with a
 period specified by the duration argument. It adjusts the intervals or drops ticks
 to make up for slow receivers. The duration d must be greater than zero; if not,
 NewTicker will panic. Stop the ticker to release associated resources.
 Stop turns off a ticker. After Stop, no more ticks will be sent. Stop does not close
 the channel, to prevent a concurrent goroutine reading from the channel from seeing
 an erroneous "tick".
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Ticker结构体持有一个按照间隔时间提交‘ticks’到channel类型变量C <- chan Time。
 2. Stop方法会关闭一个ticker。Stop执行完成后，不会再有ticks被发送，但Stop方法不会关闭channel，是为了
 防止并行的goroutine看到错误的tick。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func tick(ch <-chan time.Time) {
	for t := range ch {
		fmt.Println(t)
	}
}

func main() {
	ticker := time.NewTicker(time.Second)

	go tick(ticker.C)

	<-time.After(5 * time.Second)
	ticker.Stop()
}
