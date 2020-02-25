/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: template.HTML
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type HTML string
 ------------------------------------------------------------------------------------
 **Description:
 HTML encapsulates a known safe HTML document fragment. It should not be used
 for HTML from a third-party, or HTML with unclosed tags or comments. The
 outputs of a sound HTML sanitizer and a template escaped by this package
 are fine for use with HTML.
 Use of this type presents a security risk: the encapsulated content should
 come from a trusted source, as it will be included verbatim in the template
 output.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
