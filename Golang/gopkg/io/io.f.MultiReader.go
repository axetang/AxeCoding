/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: io
 **Element: io.MultiReader
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func MultiReader(readers ...Reader) Reader
 ------------------------------------------------------------------------------------
 **Description:
 1. MultiReader returns a Reader that's the logical concatenation of the provided
 input readers. They're read sequentially. Once all inputs have returned EOF, Read
 will return EOF. If any of the readers return a non-nil, non-EOF error, Read will
 return that error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. MultiReader函数将多个变参中的Reader成员组合起来作为一个新的Reader返回；
 2. 当以这个新的Reader作为读取源时，将自动按顺序连续读取所有组成这个新Reader的所有成员Reader；
 3. 当所有成员Reader都返回EOF时，Read操作会返回EOF；
 4. 在顺序读取中任何一个reader出错返回non-nil或non-EOF的错误，读操作会终止并返回相应错误。
*************************************************************************************/
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r4 := strings.NewReader("fourth reader\t finally\n")
	r := io.MultiReader(r1, r2, r3, r4)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
