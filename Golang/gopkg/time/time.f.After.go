/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.After
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func After(d Duration) <-chan Time
 ------------------------------------------------------------------------------------
 **Description:
 After waits for the duration to elapse and then sends the current time on the
 returned channel. It is equivalent to NewTimer(d).C. The underlying Timer is not
 recovered by the garbage collector until the timer fires. If efficiency is a concern,
 use NewTimer instead and call Timer.Stop if the timer is no longer needed.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 等待指定时间段之后将当前时间发送到返回的chan中。等价于NewTimer(d).C;
 2.
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	println("--After 1----")
	TestAfter1()
	println("--After 2----")
	//TestAfter2()
}
func TestAfter2() {
	result := make(chan int)
	go func(ch chan int) {
		time.Sleep(3 * time.Second)
		ch <- 4
	}(result)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Time out")
	case <-result:
		fmt.Println(result)
	}

}
func handle(m int) {
	fmt.Println("handle m=", m)
}

func TestAfter1() {
	c := make(chan int)
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
}
