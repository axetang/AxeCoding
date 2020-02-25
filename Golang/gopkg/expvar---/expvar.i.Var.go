/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Var
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Var interface {
     // String returns a valid JSON value for the variable.
     // Types with String methods that do not return valid JSON
     // (such as time.Time) must not be used as a Var.
     String() string
 }
 ------------------------------------------------------------------------------------
 **Description:
 Var is an abstract type for all exported variables.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Var接口是所有导出变量的抽象类型;
 2. String方法返回导出变量的一个合法的JSON值；
 3. 注意：实现了String方法的类型如果不能返回一个合法的JSON值，就不能当做Var使用。
*************************************************************************************/
package main

func main() {
}
