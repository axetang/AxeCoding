/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: path
 **Element: path.Join
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Join(elem ...string) string
 ------------------------------------------------------------------------------------
 **Description:
 Join joins any number of path elements into a single path, adding a separating
 slash if necessary. The result is Cleaned; in particular, all empty strings
 are ignored.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Join函数用于连接路径，把多个elem字符串参数链接为一个path，如果需要增加斜杠；
 2. 返回的结果是用Clean函数清理过的；
 3. 如果elem是空路径就忽略。
*************************************************************************************/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("a", "b", "c"))          // a/b/c
	fmt.Println(path.Join("a", "", "c"))           // a/c
	fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c
}
