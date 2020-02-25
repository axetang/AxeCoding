/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: log
 **Element: log.Println
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Println(v ...interface{})
 ------------------------------------------------------------------------------------
 **Description:
 Println calls Output to print to the standard logger. Arguments are handled in the
 manner of fmt.Println.
 ------------------------------------------------------------------------------------
 **要点总结:
 Go内置的Logger函数，参看Logger结构体的方法。
*************************************************************************************/
package main

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
