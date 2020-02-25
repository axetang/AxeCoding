/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.AppendBool,strconv.AppendFloat,strconv.AppendInt,strconv.AppendQuote,
 strconv.AppendQuoteRune,strconv.AppendQuoteRuneToASCII,strconv.AppendQuoteRuneToGraphic,
 strconv.AppendQuoteToASCII,strconv.AppendQuoteToGraphic,strconv.AppendUint
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func AppendBool(dst []byte, b bool) []byte
 func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
 func AppendInt(dst []byte, i int64, base int) []byte
 func AppendQuote(dst []byte, s string) []byte
 func AppendQuoteRune(dst []byte, r rune) []byte
 func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
 func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
 func AppendQuoteToASCII(dst []byte, s string) []byte
 func AppendQuoteToGraphic(dst []byte, s string) []byte
 func AppendUint(dst []byte, i uint64, base int) []byte
 ------------------------------------------------------------------------------------
 **Description:a
 AppendFloat appends the string form of the floating-point number f, as generated by
 FormatFloat, to dst and returns the extended buffer.
 AppendInt appends the string form of the integer i, as generated by FormatInt, to dst
 and returns the extended buffer.
 AppendQuote appends a double-quoted Go string literal representing s, as generated
 by Quote, to dst and returns the extended buffer.
 AppendQuoteRune appends a single-quoted Go character literal representing the rune,
 as generated by QuoteRune, to dst and returns the extended buffer.
 AppendQuoteRuneToASCII appends a single-quoted Go character literal representing the
 rune, as generated by QuoteRuneToASCII, to dst and returns the extended buffer.
 AppendQuoteRuneToGraphic appends a single-quoted Go character literal representing
 the rune, as generated by QuoteRuneToGraphic, to dst and returns the extended buffer.
 AppendQuoteToASCII appends a double-quoted Go string literal representing s, as
 generated by QuoteToASCII, to dst and returns the extended buffer.
 AppendQuoteToGraphic appends a double-quoted Go string literal representing s, as
 generated by QuoteToGraphic, to dst and returns the extended buffer.
 AppendUint appends the string form of the unsigned integer i, as generated by
 FormatUint, to dst and returns the extended buffer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. AppendBool函数向字节切片参数dst尾部写入参数b的bool值"true"或"false";
 2. AppendFloat函数向字节切片参数dst尾部写入浮点参数f依据fmt，prec和bitSize确定的值显示字符串;
 3. AppendBool函数向字节切片参数dst尾部写入参数i依据base确定的的整数值显示字符串;
 4. AppendQuote函数向字符切片参数dst尾部写入参数s（使用双引用包裹的字符串字面量）；
 5. AppendQuoteRune函数向字符切片参数dst尾部写数单引号包裹的参数r rune；
 6. AppendQuoteRuneToASCII函数向字符切片参数dst尾部写入单引号包裹的参数r rune的Ascii值；
 7. AppendQuoteRuneToGraphic函数向字符切片参数dst尾部写入单引号包裹的参数r rune的图形；
 8. AppendQuoteToASCII函数向字符切片参数dst尾部写入单引号包裹的参数s string的Ascii值；
 9. AppendQuoteToGraphic函数向字符切片参数dst尾部写入单引号包裹的参数s string的图形；
10. AppendUint函数向字节切片参数dst尾部写入参数i uint64依据base确定的的整数值显示字符串;
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--------AppendBool---------")
	s := []byte("12345")
	sn := strconv.AppendBool(s, true)
	fmt.Println(string(sn))

	fmt.Println("--------AppendFloat---------")
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

	fmt.Println("--------AppendInt---------")
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b2 := []byte("int (base 2):")
	b2 = strconv.AppendInt(b2, -42, 2)
	fmt.Println(string(b2))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	fmt.Println("--------AppendQuote---------")
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))
	fmt.Println("--------AppendQuoteRune---------")
	b = []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))
	fmt.Println("--------AppendQuoteRuneTOASCII---------")
	b = []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))
	fmt.Println("--------AppendQuoteRuneToGraphic---------")
	b = []byte("rune (Graphic):")
	b = strconv.AppendQuoteRuneToGraphic(b, '☺')
	fmt.Println(string(b))

	fmt.Println("--------AppendQuoteToASCII---------")
	b = []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	fmt.Println("--------AppendQuoteToGraphic---------")
	b = []byte("rune (Graphic):")
	b = strconv.AppendQuoteToGraphic(b, "☺")
	fmt.Println(string(b))

	fmt.Println("--------AppendUint---------")
	b10 = []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))

	b16 = []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16))
}