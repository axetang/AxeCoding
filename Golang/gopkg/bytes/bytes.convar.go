/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bytes
 **Element: convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const MinRead = 512
 var ErrTooLarge = errors.New("bytes.Buffer: too large")
 ------------------------------------------------------------------------------------
 **Description:
 MinRead is the minimum slice size passed to a Read call by Buffer.ReadFrom. As long
 as the Buffer has at least MinRead bytes beyond what is required to hold the contents
 of r, ReadFrom will not grow the underlying buffer.
 ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1.
*************************************************************************************/
package main

func main() {
}
