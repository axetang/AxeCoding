/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.UnquoteUsage
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func UnquoteUsage(flag *Flag) (name string, usage string)
------------------------------------------------------------------------------------
 **Description:
 UnquoteUsage extracts a back-quoted name from the usage string for a flag 
 and returns it and the un-quoted usage. Given "a `name` to show" it returns 
 ("name", "a name to show"). If there are no back quotes, the name is an 
 educated guess of the type of the flag's value, or the empty string if the 
 flag is boolean.
 -------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
