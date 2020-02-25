/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ioutil
 **Element: ioutil.ReadAll
 **Type:func
 ------------------------------------------------------------------------------------
 **Definition:
 func ReadAll(r io.Reader) ([]byte, error)
 ------------------------------------------------------------------------------------
 **Description:
 ReadAll reads from r until an error or EOF and returns the data it read. A
 successful call returns err == nil, not err == EOF. Because ReadAll is defined to
 read from src until EOF, it does not treat an EOF from Read as an error to be reported.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ReadAll从参数r中读取全部数据，直到EOF或者出现error；
 2. ReadAll返回值中的[]byte包含所有读取的字节内容；
 2. 成功的ReadAll返回err==nil，而不是EOF。因为ReadAll的定义是从r中读取数据直到EOF，它不将EOF看成error。
*************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("err is ", err)
	fmt.Printf("%s\n", b)
}
