/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.FormatBool,strconv.FormatFloat,strconv.FormatInt,strconv.FormatUint
 **Type: func
 **Note: This code segment was completed in Hongkong Univ. on Feb. 15, 2019.
 ------------------------------------------------------------------------------------
 **Definition:
 func FormatBool(b bool) string
 func FormatFloat(f float64, fmt byte, prec, bitSize int) string
 func FormatInt(i int64, base int) string
 func FormatUint(i uint64, base int) string
 ------------------------------------------------------------------------------------
 **Description:
 FormatBool returns "true" or "false" according to the value of b.
 FormatFloat converts the floating-point number f to a string, according to the
 format fmt and precision prec. It rounds the result assuming that the original was
 obtained from a floating-point value of bitSize bits (32 for float32, 64 for float64).
 The format fmt is one of 'b' (-ddddp±ddd, a binary exponent), 'e' (-d.dddde±dd, a
 decimal exponent), 'E' (-d.ddddE±dd, a decimal exponent), 'f' (-ddd.dddd, no
 exponent), 'g' ('e' for large exponents, 'f' otherwise), or 'G' ('E' for large
 exponents, 'f' otherwise).
 The precision prec controls the number of digits (excluding the exponent) printed
 by the 'e', 'E', 'f', 'g', and 'G' formats. For 'e', 'E', and 'f' it is the number
 of digits after the decimal point. For 'g' and 'G' it is the maximum number of
 significant digits (trailing zeros are removed). The special precision -1 uses the
 smallest number of digits necessary such that ParseFloat will return f exactly.
 FormatInt returns the string representation of i in the given base, for 2<=base<=36.
 The result uses the lower-case letters 'a' to 'z' for digit values >= 10.
 FormatUint returns the string representation of i in the given base, for 2<=base<=36.
 The result uses the lower-case letters 'a' to 'z' for digit values >= 10.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FormatBool依照参数b bool的值来返回"true"或者"false"字符串；
 2. FormatFloat依照参数fmt、prec、bitSize来格式化和取整并转化浮点数f float64为字符串返回，
 具体如下：
	1）fmt byte参数可以选取的值为：'b','e','E','f','g','G';
	2) prec int参数可以选取的值为：'e','E','f','g','G'，'-1';
 3. FormatInt返回按照参数base确定的进制编码后的有符号整数i int64的字符串；
 4. FormatInt返回按照参数base确定的进制编码后的无符号整数i int64的字符串；
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := true
	s := strconv.FormatBool(v)
	fmt.Printf("%T, %v\n", s, s)
	println()

	v1 := 3.1415926535
	s32 := strconv.FormatFloat(v1, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
	s64 := strconv.FormatFloat(v1, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
	println()

	v2 := int64(-42)
	s10 := strconv.FormatInt(v2, 10)
	fmt.Printf("%T, %v\n", s10, s10)
	s16 := strconv.FormatInt(v2, 16)
	fmt.Printf("%T, %v\n", s16, s16)
	println()

	v3 := uint64(42)
	ss10 := strconv.FormatUint(v3, 10)
	fmt.Printf("%T, %v\n", ss10, ss10)
	ss16 := strconv.FormatUint(v3, 16)
	fmt.Printf("%T, %v\n", ss16, ss16)
}
