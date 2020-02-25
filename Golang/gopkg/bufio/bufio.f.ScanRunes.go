/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: bufio
 **Element: bufio.ScanRunes
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
 ------------------------------------------------------------------------------------
 **Description:
 ScanRunes is a split function for a Scanner that returns each UTF-8-encoded rune
 as a token. The sequence of runes returned is equivalent to that from a range
 loop over the input as a string, which means that erroneous UTF-8 encodings
 translate to U+FFFD = "\xef\xbf\xbd". Because of the Scan interface, this makes
 it impossible for the client to distinguish correctly encoded replacement runes
 from encoding errors.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. 在bufio包中预定义了一些split函数，也就是说，在Scanner结构中的split字段，可以通过这些预定义
 的split赋值，同时Scanner类型的Split方法也可以接收这些预定义函数作为参数。所以，我们可以说这些
 预定义split函数都是SplitFunc类型的实例。这些函数包括：ScanBytes、ScanRunes、ScanWords和
 ScanLines, 由于都是SplitFunc的实例，自然这些函数的签名都和SplitFunc一样;
 1.ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个utf-8编码的unicode
 码值作为一个token返回。 本函数返回的rune序列和range一个字符串的输出rune序列相同。 错误的
 utf-8编码会翻译为U+FFFD = "\xef\xbf\xbd"， 但只会消耗一个字节。调用者无法区分正确编码的
 rune和错误编码的rune。
*************************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "我们的祖国祥花园\n哈哈 我是谁、\n 作者By Axe."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
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
