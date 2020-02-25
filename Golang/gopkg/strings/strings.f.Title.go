/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: strings
**Element: strings.Title
**Type: func
------------------------------------------------------------------------------------
**Definition:
func Title(s string) string
------------------------------------------------------------------------------------
**Description:
Title returns a copy of the string s with all Unicode letters that begin words
mapped to their title case.
------------------------------------------------------------------------------------
**要点总结:
Title方法将字符串s按照首字母大写的方式修改后返回；
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("herroyalhighness"))
}
