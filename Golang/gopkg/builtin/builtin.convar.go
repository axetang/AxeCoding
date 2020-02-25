/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    true  = 0 == 0 // Untyped bool.
    false = 0 != 0 // Untyped bool.
 )
 const iota = 0 // Untyped int.
 ------------------------------------------------------------------------------------
 **Description:
 true and false are the two untyped boolean values.
 iota is a predeclared identifier representing the untyped integer ordinal number of
 the current const specification in a (usually parenthesized) const declaration. It is
 zero-indexed.
 ------------------------------------------------------------------------------------
 **要点总结:

*************************************************************************************/
package main

func main() {
}
