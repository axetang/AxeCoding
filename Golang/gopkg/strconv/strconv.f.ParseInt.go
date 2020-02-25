/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.ParseInt
 **Type: func
 **Note: 此段程序于2019年2月15日在香港大学完成。
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseInt(s string, base int, bitSize int) (i int64, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64)
 and returns the corresponding value i.
 If base == 0, the base is implied by the string's prefix: base 16 for "0x", base 8 for
 "0", and base 10 otherwise. For bases 1, below 0 or above 36 an error is returned.
 The bitSize argument specifies the integer type that the result must fit into. Bit
 sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64. For a
 bitSize below 0 or above 64 an error is returned.
 The errors that ParseInt returns have concrete type *NumError and include
 err.Num = s. If s is empty or contains invalid digits, err.Err = ErrSyntax and the
 returned value is 0; if the value corresponding to s cannot be represented by a
 signed integer of the given size, err.Err = ErrRange and the returned value is the
 maximum magnitude integer of the appropriate bitSize and sign.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseInt将字符串s按照进制参数base(0~32)进行转化，转化的值范围使用bitSize(0~64)参数确定，即2
 的(bitsize-1)次方，返回值i类型为int64；
 2. bitSize取值为0，8，16，32，64,对应int,int8,int16,int32,int64；
 3. 注意，ParseInt转化字符串为有符号整数，即二进制首位为符号位；
 4. 如果转化值超过bitSize规定的最大值，参数返回错误：value out of range。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v2 := "-511"
	if s, err := strconv.ParseInt(v2, 10, 10); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println("error:", err)
	}

	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
