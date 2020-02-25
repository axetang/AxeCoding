/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strconv
 **Element: strconv.NumError
 **Type: struct
 **Note: this code segment was done in Hongkong Disney Hotel on Feb. 16, 2019.
 ------------------------------------------------------------------------------------
 **Definition:
 type NumError struct {
    Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
    Num  string // the input
    Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
 }
 func (e *NumError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 A NumError records a failed conversion.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NumError结构体记录了一次失败的转换的相关信息；
 2. strconv包中函数返回的error接口类型的实际类型都是*strconv.NumError；
 3. Go语言类型转换的一个方法叫做断言。断言的语法是 x.(T)，x指的是类型为interface{}的变量，T是
 我们断言它可能是的类型。该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，
 若为true则表示断言成功，false则表示断言失败，如下程序示例。
 4. Error方法具体实现如下：
 func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing " + Quote(e.Num) + ": " +
	e.Err.Error()
 }
*************************************************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "Not a number"
	_, err := strconv.ParseFloat(str, 64)
	fmt.Printf("err type is %T, value is %v\n", err, err)
	if err != nil {
		e := err.(*strconv.NumError)
		fmt.Printf("e type: %T\n", e)
		fmt.Println("Func:", e.Func)
		fmt.Println("Num:", e.Num)
		fmt.Println("Err:", e.Err)
		fmt.Println("e.Error():", e.Error())
		fmt.Println(err)
	}
}
