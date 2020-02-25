/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: tempalte.Must
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Must(t *Template, err error) *Template
 ------------------------------------------------------------------------------------
 **Description:
 Must is a helper that wraps a call to a function returning (*Template, error)
 and panics if the error is non-nil. It is intended for use in variable
 initializations such as
 var t = template.Must(template.New("name").Parse("html"))
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
