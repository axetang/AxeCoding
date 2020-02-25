/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: crc32
 **Element: crc32.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    // IEEE is by far and away the most common CRC-32 polynomial.
    // Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
    IEEE = 0xedb88320

    // Castagnoli's polynomial, used in iSCSI.
    // Has better error detection characteristics than IEEE.
    // https://dx.doi.org/10.1109/26.231911
    Castagnoli = 0x82f63b78

    // Koopman's polynomial.
    // Also has better error detection characteristics than IEEE.
    // https://dx.doi.org/10.1109/DSN.2002.1028931
    Koopman = 0xeb31d82e
 )
 Predefined polynomials.
 const Size = 4
 The size of a CRC-32 checksum in bytes.
 var IEEETable = simpleMakeTable(IEEE)
 IEEETable is the table for the IEEE polynomial.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
