/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.ScanBytes
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ScanBytes is a split function for a Scanner that returns each byte as a token.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 在 bufio 包中预定义了一些split函数，也就是说，在Scanner结构中的split字段，可以通过这些预定义
 的split赋值，同时Scanner类型的Split方法也可以接收这些预定义函数作为参数。所以，我们可以说这些
 预定义split函数都是SplitFunc类型的实例。这些函数包括：ScanBytes、ScanRunes、ScanWords和
 ScanLines, 由于都是SplitFunc的实例，自然这些函数的签名都和SplitFunc一样;
 1. ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token
 返回。
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
	scanner.Split(bufio.ScanBytes)
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
