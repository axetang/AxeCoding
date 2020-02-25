/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.ParseUint
 **Type: func
 **Note: 此段程序于2019年2月15日在香港大学完成。
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseUint(s string, base int, bitSize int) (uint64, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseUint is like ParseInt but for unsigned numbers.
 ------------------------------------------------------------------------------------
 **要点总结:
 ParseUint和ParseInt类似，唯一不同是将字符串转化为无符号整数。
 1. ParseInt将字符串s按照进制参数base(0~32)进行转化，转化的值范围使用bitSize(0~64)参数确定，即
 0-2的(bitsize)次方，返回值i类型为uint64；
 2. bitSize取值为0，8，16，32，64,对应uint,uint8,uint16,uint32,uint64；
 3. 注意，ParseInt转化字符串为无符号整数；
 4. 如果转化值超过bitSize规定的最大值，参数返回错误：value out of range。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
