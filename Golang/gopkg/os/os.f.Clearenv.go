/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Clearenv
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Clearenv()
 ------------------------------------------------------------------------------------
 **Description:
 Clearenv deletes all environment variables.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Clearenv函数清空所有环境变量;
 2. 并不是清楚shell中的变量，而是在该程序的执行期立即清楚所有环境变量。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The path is: %+v\n", os.Environ())
	os.Clearenv()
	fmt.Printf("Now the path is: %+v\n", os.Environ())
}
