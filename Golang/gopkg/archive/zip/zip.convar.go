/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: zip
 **Element: zip.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
     Store   uint16 = 0 // no compression
     Deflate uint16 = 8 // DEFLATE compressed
 )
 Compression methods.
 var (
     ErrFormat    = errors.New("zip: not a valid zip file")
     ErrAlgorithm = errors.New("zip: unsupported compression algorithm")
     ErrChecksum  = errors.New("zip: checksum error")
 )
 ------------------------------------------------------------------------------------
 **Description:
 NA
 ------------------------------------------------------------------------------------
 **要点总结:
 NA
*************************************************************************************/
package main

func main() {
}
