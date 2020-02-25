/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Tick
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Tick(d Duration) <-chan Time
 ------------------------------------------------------------------------------------
 **Description:
 Tick is a convenience wrapper for NewTicker providing access to the ticking channel
 only. While Tick is useful for clients that have no need to shut down the Ticker, be
 aware that without a way to shut it down the underlying Ticker cannot be recovered
 by the garbage collector; it "leaks". Unlike NewTicker, Tick will return nil
 if d <= 0.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Tick是对NewTicker更易用的包装，提供了对定时器channel的访问；
 2. 通常用于没必要停止定时器的客户端；
 3. 此函数返回一个Time类型的chan；
 4. 此函数运行机制还要在深入了解下后更新。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\n", now)
	}
}
