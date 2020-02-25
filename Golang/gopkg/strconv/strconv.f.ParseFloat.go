/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.ParseFloat
 **Type: func
 **Note: 此段程序于2019年2月15日在香港大学完成。
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseFloat(s string, bitSize int) (float64, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseFloat converts the string s to a floating-point number with the precision
 specified by bitSize: 32 for float32, or 64 for float64. When bitSize=32, the
 result still has type float64, but it will be convertible to float32 without
 changing its value.
	 If s is well-formed and near a valid floating point number, ParseFloat returns
	 the nearest floating point number rounded using IEEE754 unbiased rounding.
	 The errors that ParseFloat returns have concrete type *NumError and include
	 err.Num = s.
 	 If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.
	 If s is syntactically well-formed but is more than 1/2 ULP away from the largest
	 floating point number of the given size, ParseFloat returns f = ±Inf,
	 err.Err = ErrRange.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseFloat函数将字符串s按照参数bitSize转化为一个float64返回；
 2. bitSize取值：32代表float32的范围；64代表float64的范围；
 3. s如果无法正确转化，则返回错误。
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//32位的转化结果需要分析下。
	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
