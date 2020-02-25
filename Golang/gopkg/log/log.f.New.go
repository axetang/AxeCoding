/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: log
 **Element: log.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(out io.Writer, prefix string, flag int) *Logger
 ------------------------------------------------------------------------------------
 **Description:
 New creates a new Logger. The out variable sets the destination to which log data
 will be written. The prefix appears at the beginning of each generated log line.
 The flag argument defines the logging properties.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. New函数创建一个Logger结构体*Logger；
 2. out io.Writer设定了log数据写入目标；
 3. prefix string字符串定义了每个log信息的前缀字符串；
 4. flag int定义了输入出日志的文件的展示方式，由log包中的常数定义如下：
  const (
    // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
    // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
 )
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {

	fmt.Println("-----New&Print-----")
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Llongfile)
	)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)
}
