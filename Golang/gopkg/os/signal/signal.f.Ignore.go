/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: signal
 **Element: signal.Ignore
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Ignore(sig ...os.Signal)
 ------------------------------------------------------------------------------------
 **Description:
 Ignore causes the provided signals to be ignored. If they are received by the
 program, nothing will happen. Ignore undoes the effect of any prior calls to
 Notify for the provided signals. If no signals are provided, all incoming signals
 will be ignored.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
