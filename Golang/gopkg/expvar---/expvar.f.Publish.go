/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Publish
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Publish(name string, v Var)
 ------------------------------------------------------------------------------------
 **Description:
 Publish declares a named exported variable. This should be called from a
 package's init function when it creates its Vars. If the name is already
 registered then this will log.Panic.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Publish声明一个导出变量。必须在init函数里调用。如果name已经被注册，会调用log.Panic。
*************************************************************************************/
package main

func main() {
}
