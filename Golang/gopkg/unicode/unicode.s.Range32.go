/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.Range32
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Range32 struct {
     Lo     uint32
     Hi     uint32
     Stride uint32
 }
 ------------------------------------------------------------------------------------
 **Description:
 Range32 represents of a range of Unicode code points and is used when one or more
 of the values will not fit in 16 bits. The range runs from Lo to Hi inclusive and
 has the specified stride. Lo and Hi must always be >= 1<<16.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Range32表示一个32位的Unicode码点范围，范围从Lo到Hi，具有指定的步长，Lo和Hi必须都大于16位。
*************************************************************************************/
package main

func main() {
}
