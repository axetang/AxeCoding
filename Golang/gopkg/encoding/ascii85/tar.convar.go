/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: tar
 **Element: tar.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
     // Type '0' indicates a regular file.
     TypeReg  = '0'
     TypeRegA = '\x00' // Deprecated: Use TypeReg instead.

     // Type '1' to '6' are header-only flags and may not have a data body.
     TypeLink    = '1' // Hard link
     TypeSymlink = '2' // Symbolic link
     TypeChar    = '3' // Character device node
     TypeBlock   = '4' // Block device node
     TypeDir     = '5' // Directory
     TypeFifo    = '6' // FIFO node

     // Type '7' is reserved.
     TypeCont = '7'

     // Type 'x' is used by the PAX format to store key-value records that
     // are only relevant to the next file.
     // This package transparently handles these types.
     TypeXHeader = 'x'

     // Type 'g' is used by the PAX format to store key-value records that
     // are relevant to all subsequent files.
     // This package only supports parsing and composing such headers,
     // but does not currently support persisting the global state across files.
     TypeXGlobalHeader = 'g'

     // Type 'S' indicates a sparse file in the GNU format.
     TypeGNUSparse = 'S'

     // Types 'L' and 'K' are used by the GNU format for a meta file
     // used to store the path or link name for the next file.
     // This package transparently handles these types.
     TypeGNULongName = 'L'
     TypeGNULongLink = 'K'
 )
 Type flags for Header.Typeflag.
 var (
     ErrHeader          = errors.New("archive/tar: invalid tar header")
     ErrWriteTooLong    = errors.New("archive/tar: write too long")
     ErrFieldTooLong    = errors.New("archive/tar: header field too long")
     ErrWriteAfterClose = errors.New("archive/tar: write after close")
 )
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
