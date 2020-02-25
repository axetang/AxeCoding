/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    Separator     = os.PathSeparator
    ListSeparator = os.PathListSeparator
 )
 var ErrBadPattern = errors.New("syntax error in pattern")
 var SkipDir = errors.New("skip this directory")
 ------------------------------------------------------------------------------------
 **Description:
 ErrBadPattern indicates a pattern was malformed.
 SkipDir is used as a return value from WalkFuncs to indicate that the
 directory named in the call is to be skipped. It is not returned as an
 error by any function.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
