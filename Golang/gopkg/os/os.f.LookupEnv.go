/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.LookupEnv
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupEnv(key string) (string, bool)
 ------------------------------------------------------------------------------------
 **Description:
 LookupEnv retrieves the value of the environment variable named by the key. If the
 variable is present in the environment the value (which may be empty) is returned
 and the boolean is true. Otherwise the returned value will be empty and the boolean
 will be false.
 ------------------------------------------------------------------------------------
 **要点总结:s
 1. LookupEnv函数在环境变量中搜索变量key的值，如果找到，返回key的值及true；
 2. 如果找不到，返回空字符串及false。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	s, b := os.LookupEnv("LANG")
	fmt.Printf("%s, %t\n", s, b)
	s, b = os.LookupEnv("NG")
	fmt.Printf("%s, %t\n", s, b)
}
