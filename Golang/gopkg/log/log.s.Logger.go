/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: log
 **Element: log.Logger
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // A Logger represents an active logging object that generates lines of
 // output to an io.Writer. Each logging operation makes a single call to
 // the Writer's Write method. A Logger can be used simultaneously from
 // multiple goroutines; it guarantees to serialize access to the Writer.
 type Logger struct {
    // contains filtered or unexported fields
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
 }
 func (l *Logger) Fatal(v ...interface{})
 func (l *Logger) Fatalf(format string, v ...interface{})
 func (l *Logger) Fatalln(v ...interface{})
 func (l *Logger) Flags() int
 func (l *Logger) Output(calldepth int, s string) error
 func (l *Logger) Panic(v ...interface{})
 func (l *Logger) Panicf(format string, v ...interface{})
 func (l *Logger) Panicln(v ...interface{})
 func (l *Logger) Prefix() string
 func (l *Logger) Print(v ...interface{})
 func (l *Logger) Printf(format string, v ...interface{})
 func (l *Logger) Println(v ...interface{})
 func (l *Logger) SetFlags(flag int)
 func (l *Logger) SetOutput(w io.Writer)
 func (l *Logger) SetPrefix(prefix string)
 func (l *Logger) Writer() io.Writer
 ------------------------------------------------------------------------------------
 **Description:
 A Logger represents an active logging object that generates lines of output to an
 io.Writer. Each logging operation makes a single call to the Writer's Write method.
 A Logger can be used simultaneously from multiple goroutines; it guarantees to
 serialize access to the Writer.
 Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
 Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
 Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
 Flags returns the output flags for the logger.
 Output writes the output for a logging event. The string s contains the text to
 print after the prefix specified by the flags of the Logger. A newline is appended
 if the last character of s is not already a newline. Calldepth is used to recover
 the PC and is provided for generality, although at the moment on all pre-defined
 paths it will be 2.
 Panic is equivalent to l.Print() followed by a call to panic().
 Panicf is equivalent to l.Printf() followed by a call to panic().
 Panicln is equivalent to l.Println() followed by a call to panic().
 Prefix returns the output prefix for the logger.
 Print calls l.Output to print to the logger. Arguments are handled in the manner of
 fmt.Print.
 Printf calls l.Output to print to the logger. Arguments are handled in the manner
 of fmt.Printf.
 Println calls l.Output to print to the logger. Arguments are handled in the manner
 of fmt.Println.
 SetFlags sets the output flags for the logger.
 SetOutput sets the output destination for the logger.
 SetPrefix sets the output prefix for the logger.
 Writer returns the output destination for the logger.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Logger类型表示一个活动状态的记录日志的对象，它会生成一行行的输出写入一个io.Writer接口。每一条
 日志操作会调用一次io.Writer接口的Write方法。Logger类型的对象可以被多个线程安全的同时使用，它会保
 证对io.Writer接口的顺序访问；
 1. Fatal等价于{l.Print(v...); os.Exit(1)};
 2. Fatalf等价于{l.Printf(v...); os.Exit(1)};
 3. Fatalln等价于{l.Println(v...); os.Exit(1)}；
 4. Flags方法返回当前Logger的Flag常数设置值，具体参见convar常数定义，常数间可以通过|组合；
 5. Output方法写入输出一次日志事件。参数s包含在Logger根据选项生成的前缀之后要打印的文本。如果s末尾
 没有换行会添加换行符。calldepth用于恢复PC，出于一般性而提供，但目前在所有预定义的路径上它的值都为2；
 6. Panic方法等价于{l.Print(v...); panic(...)}；
 7. Panicf方法等价于{l.Printf(v...); panic(...)}；
 8. Panicln方法等价于{l.Println(v...); panic(...)}；
 9. Prefix方法Prefix返回logger的输出前缀；
 10. Print方法调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Print相同的方法处理;
 11. Printf方法调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Printf相同的方法处理;
 12. Println方法调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Print相同的方法处理;
 13. SetFlags方法设置当前Logger的输出选项，以flag int值的方式来设定；
 14. SetOutput方法设置当前Logger的输出目标w io.Writer；
 15. SetPrefix方法设置logger的输出前缀；
 16. Writer方法返回当前Logger的输出目标，可以使用如下方法获得Stdout输出目标：
 	l.Writer().(*os.File).Name()
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("-----New&Print-----")
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)

	fmt.Println("-----Flags-----")
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	l.Print("Hello,l")
	fmt.Println("logger l's flags :", l.Flags())

	fmt.Printf("Ldate   %d      %8b\n", log.Ldate, log.Ldate)
	fmt.Printf("Ltime   %d      %b\n", log.Ltime, log.Ltime)
	fmt.Printf("Lmicroseconds   %d      %b\n", log.Lmicroseconds, log.Lmicroseconds)
	fmt.Printf("Llongfile   %d      %b\n", log.Llongfile, log.Llongfile)
	fmt.Printf("Lshortfile   %d      %b\n", log.Lshortfile, log.Lshortfile)
	fmt.Printf("LUTC   %d      %b\n", log.LUTC, log.LUTC)
	fmt.Printf("LstdFlags   %d      %b\n", log.LstdFlags, log.LstdFlags)

	fmt.Println("-----Flags-SetFlags，Prefix-SetPrefix,SetOutput-Writer, Output-----")
	l = log.New(os.Stdout, "Axe:", log.LstdFlags)
	l.Println(l.Flags(), l.Prefix())
	l.SetFlags(log.Ldate)
	l.SetPrefix("Tang:")
	l.SetOutput(os.Stdout)
	l.Println(l.Flags(), l.Prefix(), l.Writer().(*os.File).Name())
	l.Output(2, "this is a log by Output()")

	fmt.Println("-----Print,Printf,Println-----")
	l.Print("this is ", "a log")
	l.Println("this is a log", 555)
	l.Printf("this is %d log", 2)

	fmt.Println("-----Fatal,Fatalf,Fatalln-----")
	//l.Fatal("this is a fatal error!")
	//l.Fatalln("this is a fatal error!")
	//l.Fatalf("this is the %dth fatal error!", 5)

	fmt.Println("-----Panic,Panicf,Panicln-----")
	//l.Panic("this is a panic by axe!")
	//l.Panicln("this is a panic by axe tang!")
	//l.Panicf("this is a panic by axe tang %d!", 5)
}
