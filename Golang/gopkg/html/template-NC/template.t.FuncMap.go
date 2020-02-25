/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.FuncMap
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type FuncMap map[string]interface{}
 ------------------------------------------------------------------------------------
 **Description:
 FuncMap is the type of the map defining the mapping from names to functions.
 Each function must have either a single return value, or two return values
 of which the second has type error. In that case, if the second (error)
 argument evaluates to non-nil during execution, execution terminates and
 Execute returns that error. FuncMap has the same base type as FuncMap in
 "text/template", copied here so clients need not import "text/template".
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
