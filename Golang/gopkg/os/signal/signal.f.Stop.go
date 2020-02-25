/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: signal
 **Element: signal.Stop
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Stop(c chan<- os.Signal)
 ------------------------------------------------------------------------------------
 **Description:
 Stop causes package signal to stop relaying incoming signals to c. It undoes the
 effect of all prior calls to Notify using c. When Stop returns, it is guaranteed
 that c will receive no more signals.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
