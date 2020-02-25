/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.FileListener
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FileListener(f *os.File) (ln Listener, err error)
 ------------------------------------------------------------------------------------
 **Description:
 FileListener returns a copy of the network listener corresponding to the open file
 f. It is the caller's responsibility to close ln when finished. Closing ln does
 not affect f, and closing f does not affect ln.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
