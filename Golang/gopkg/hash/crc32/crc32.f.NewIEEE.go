/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: crc32
 **Element: crc32.NewIEEE
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewIEEE() hash.Hash32
 ------------------------------------------------------------------------------------
 **Description:
 NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using the IEEE
 polynomial. Its Sum method will lay the value out in big-endian byte order.
 The returned Hash32 also implements encoding.BinaryMarshaler and
 encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
