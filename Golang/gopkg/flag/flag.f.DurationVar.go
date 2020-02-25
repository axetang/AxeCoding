/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.DurationVar
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func DurationVar(p *time.Duration, name string, value time.Duration, usage
 string)
------------------------------------------------------------------------------------
 **Description:
 DurationVar defines a time.Duration flag with specified name, default value,
 and usage string. The argument p points to a time.Duration variable in which
 to store the value of the flag. The flag accepts a value acceptable to
 time.ParseDuration.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. DurationVar定义一个time.Duration flag;
 2. 函数参数指定了flag的名字，缺省值和Usage；
 3. 第一个参数p *time.Duration指向一个time.Duration变量，它存储着flag的值；
 4. 在命令行输入flag的值时，必须符合time.ParseDuration函数要求，即时间数值后面要带上时间
 单位，如：1m,1s,1h,1ns,1ms,etc.
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"time"
)

//必须在DurationVar函数使用flag变量前定义du
var du time.Duration

func init() {
	flag.DurationVar(&du, "du", time.Second*2, "-du 2000")
}

func main() {
	flag.Parse()
	fmt.Println("du:", du)
	fmt.Println("du:", du.Hours(), du.Minutes(), du.Seconds())

}
