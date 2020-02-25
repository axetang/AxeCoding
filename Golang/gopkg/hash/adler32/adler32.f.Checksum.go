/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: adler32
 **Element: adler32.Checksum
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Checksum(data []byte) uint32
 ------------------------------------------------------------------------------------
 **Description:
 Checksum returns the Adler-32 checksum of data.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	d := []byte("I am axe tang in Nanjing!")
	cs := adler32.Checksum(d)
	fmt.Printf("Checksum result: %d %b %o %x %X %U\n", cs, cs, cs, cs, cs, cs)
}
