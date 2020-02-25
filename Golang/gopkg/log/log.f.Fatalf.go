/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: log
 **Element: log.Fatalf
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fatalf(format string, v ...interface{})
 ------------------------------------------------------------------------------------
 **Description:
 Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
 ------------------------------------------------------------------------------------
 **要点总结:
 Go内置的Logger函数，参看Logger结构体的方法。
*************************************************************************************/
package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("Axe:")
	log.SetOutput(os.Stdout)
	log.SetFlags(19)
	log.Print("This is the first log.")
	log.Println("Flags()=", log.Flags())
	log.Printf("Prefix is %s\n", log.Prefix())
	log.Output(2, "this is log by Output.")
	//log.Fatal("this is a fatal error!")
	//log.Fatalln("this is a fatal error by Fatalln!")
	//log.Fatalf("this is %dth fatal error!", 5)

	//log.Panic("this is a panic by Panic!")
	//log.Panic("this is a panic by Panic!")
	//log.Panic("this is a panic by Panic!")

}
