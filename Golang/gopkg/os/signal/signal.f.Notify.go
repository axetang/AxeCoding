/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: signal
 **Element: signal.Notify
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Notify(c chan<- os.Signal, sig ...os.Signal)
 ------------------------------------------------------------------------------------
 **Description:
 Notify causes package signal to relay incoming signals to c. If no signals are
 provided, all incoming signals will be relayed to c. Otherwise, just the provided
 signals will.
 Package signal will not block sending to c: the caller must ensure that c has
 sufficient buffer space to keep up with the expected signal rate. For a channel
 used for notification of just one signal value, a buffer of size 1 is sufficient.
 It is allowed to call Notify multiple times with the same channel: each call
 expands the set of signals sent to that channel. The only way to remove signals
 from the set is to call Stop.
 It is allowed to call Notify multiple times with different channels and the same
 signals: each channel receives copies of incoming signals independently.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
