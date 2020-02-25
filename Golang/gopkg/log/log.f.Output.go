/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: log
 **Element: log.Output
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Output(calldepth int, s string) error
 ------------------------------------------------------------------------------------
 **Description:
 Output writes the output for a logging event. The string s contains the text to
 print after the prefix specified by the flags of the Logger. A newline is appended
 if the last character of s is not already a newline. Calldepth is the count of the
 number of frames to skip when computing the file name and line number if Llongfile
 or Lshortfile is set; a value of 1 will print the details for the caller of Output.
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
