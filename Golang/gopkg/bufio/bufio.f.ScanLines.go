/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.ScanLines
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ScanLines is a split function for a Scanner that returns each line of text,
 stripped of any trailing end-of-line marker. The returned line may be empty.
 The end-of-line marker is one optional carriage return followed by one mandatory
 newline. In regular expression notation, it is `\r?\n`. The last non-empty line
 of input will be returned even if it has no newline.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 在bufio包中预定义了一些split函数，也就是说，在Scanner结构中的split字段，可以通过这些预定义
 的split赋值，同时Scanner类型的Split方法也可以接收这些预定义函数作为参数。所以，我们可以说这些
 预定义split函数都是SplitFunc类型的实例。这些函数包括：ScanBytes、ScanRunes、ScanWords和
 ScanLines, 由于都是SplitFunc的实例，自然这些函数的签名都和SplitFunc一样;
 1. ScanLines是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每一行文本去掉末尾的换行
 标记作为一个token返回。 返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。
 最后一行即使没有换行符也会作为一个token返回。
*************************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "This is The Golang Standard Library.\nWelcome you!\nBy Axe."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}
