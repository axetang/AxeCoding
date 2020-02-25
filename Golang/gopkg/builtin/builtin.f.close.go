/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.close
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func close(c chan<- Type)
 ------------------------------------------------------------------------------------
 **Description:
 The close built-in function closes a channel, which must be either bidirectional or
 send-only. It should be executed only by the sender, never the receiver, and has
 the effect of shutting down the channel after the last sent value is received.
 After the last value has been received from a closed channel c, any receive from
 c will succeed without blocking, returning the zero value for the channel element.
 The form x, ok := <-c will also set ok to false for a closed channel.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. close内建函数关闭信道，该信道必须为双向的或只发送的;
 2. 它只能由发送者执行，不应由接收者执行，其效果是在最后发送的值被接收后停止该信道;
 3. 已关闭的信道c中最后一个值被接收后，任何从信道c的接收操作都会无阻塞成功，关闭的信道会返回该信道
 元素类型的零值。对于已关闭的信道
 		x, ok := <-c
 还会将 ok 置为 false。
*************************************************************************************/
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 因为使用了range调用,必须 close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // range 会不断从 channel 接收值，直到它被关闭。
		fmt.Println(i)
	}
}
