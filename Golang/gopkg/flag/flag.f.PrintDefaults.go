/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.PrintDefaults
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func PrintDefaults()
------------------------------------------------------------------------------------
 **Description:
 PrintDefaults prints, to standard error unless configured otherwise, a usage
 message showing the default settings of all defined command-line flags. For
 an integer valued flag x, the default output has the form
 -x int
 	usage-message-for-x (default 7)
 The usage message will appear on a separate line for anything but a bool
 flag with a one-byte name. For bool flags, the type is omitted and if the
 flag name is one byte the usage message appears on the same line. The
 parenthetical default is omitted if the default is the zero value for the
 type. The listed type, here int, can be changed by placing a back-quoted
 name in the flag's usage string; the first such item in the message is taken
 to be a parameter name to show in the message and the back quotes are
 stripped from the message when displayed. For instance, given
 flag.String("I", "", "search `directory` for include files")
 the output will be
 -I directory
 	search directory for include files.
 To change the destination for flag messages, call CommandLine.SetOutput.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. PrintDefaults函数打印所有flag的；
 2. 对于一个integer类型的flag x，执行PrintDefaults函数后，其打印格式为
 	 -x int
 		usage-message-for-x (default 7)

*************************************************************************************/
package main

import (
	"flag"
	"fmt"
)

var strFlag = flag.String("s", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	i := flag.NFlag()
	fmt.Println("NFlag():", i)
	//执行如下命令验证程序执行结果：
	//go run flag.f.NArg.go -s=1 -bool=T a b c

	flag.PrintDefaults()
}
