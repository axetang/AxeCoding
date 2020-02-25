/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Func
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Func func() interface{}
 func (f Func) String() string
 func (f Func) Value() interface{}
 ------------------------------------------------------------------------------------
 **Description:
 Func implements Var by calling the function and formatting the returned value
 using JSON.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
