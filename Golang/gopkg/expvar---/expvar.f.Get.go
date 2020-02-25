/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Get
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Get(name string) Var
 ------------------------------------------------------------------------------------
 **Description:
 Get retrieves a named exported variable. It returns nil if the name has not
 been registered.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
)

func main() {
	iInt := expvar.NewInt("iIntName")
	iInt.Set(100)
	fmt.Println(iInt.Value(), iInt.String())
	iInt.Add(200)
	fmt.Println(iInt.Value(), iInt.String())

	sStr := expvar.NewString("sStr")
	sStr.Set("Axe Tang")
	fmt.Println(sStr.Value())

	iVar := expvar.Get("iIntName")
	fmt.Println(iVar.(*expvar.Int).Value())
	iVar = expvar.Get("iInt")
	fmt.Println(iVar)
}
