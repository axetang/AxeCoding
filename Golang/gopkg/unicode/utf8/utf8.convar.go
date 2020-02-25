/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: utf8
 **Element: utf8.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
    RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
    MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
    UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
 )
 ------------------------------------------------------------------------------------
 **Description:
 Numbers fundamental to the encoding.
 ------------------------------------------------------------------------------------
 **要点总结:
 1.
*************************************************************************************/
package main

func main() {
}
