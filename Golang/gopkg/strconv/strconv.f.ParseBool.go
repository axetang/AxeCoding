/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.ParseBool
 **Type: func
 **Note: 此段程序于2019年2月15日在香港大学完成。
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseBool(str string) (bool, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseBool returns the boolean value represented by the string. It accepts 1, t, T,
 TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseBool依据参数str的值返回它代表的bool值：true或false；
 2. 当str取值为1，t，T，TRUE，true，True，0，f，F，False，false，False时可以正确返回；否则，
 返回错误。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v = "tRue"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println("err", err)
	}

}
