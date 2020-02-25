/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Duration
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Duration(name string, value time.Duration, usage string) *time.Duration
------------------------------------------------------------------------------------
 **Description:
 Duration defines a time.Duration flag with specified name, default value,
 and usage string. The return value is the address of a time.Duration
 variable that stores the value of the flag. The flag accepts a value
 acceptable to time.ParseDuration.
 -------------------------------------------------------------------------------------
 **要点总结:
 1. DurationVar定义一个time.Duration flag;
 2. 函数参数指定了flag的名字，缺省值和Usage；
 3. 返回值是这个time.Duration变量的地址，它指向flag的值；
 4. 在命令行输入flag的值的字面量时，要给数字加上时间单位，如：1s,1000ns,1h，1m,etc。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"time"
)

//必须在Duration函数使用flag变量前定义全局变量du
var du = flag.Duration("du", time.Second*2, "-du 2000")

func main() {
	flag.Parse()
	fmt.Println("duFlag:", du)
	fmt.Println("duFlag:", du.Hours(), du.Minutes(), du.Seconds())

}
