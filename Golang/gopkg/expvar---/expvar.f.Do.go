/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Do
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Do(f func(KeyValue))
 ------------------------------------------------------------------------------------
 **Description:
 Do calls f for each exported variable. The global variable map is locked during
 the iteration, but existing entries may be concurrently updated.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
