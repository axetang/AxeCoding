/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: tempalte.ParseGlob
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseGlob(pattern string) (*Template, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseGlob creates a new Template and parses the template definitions from
 the files identified by the pattern, which must match at least one file.
 The returned template will have the (base) name and (parsed) contents of
 the first file matched by the pattern. ParseGlob is equivalent to calling
 ParseFiles with the list of files matched by the pattern.
 When parsing multiple files with the same name in different directories,
 the last one mentioned will be the one that results.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
