/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strings
 **Element: strings.Replacer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Replacer struct {
	// contains filtered or unexported fields
	once   sync.Once // guards buildOnce method
	r      replacer
	oldnew []string
 }
 func NewReplacer(oldnew ...string) *Replacer
 func (r *Replacer) Replace(s string) string
 func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
 ------------------------------------------------------------------------------------
 **Description:
 Replacer replaces a list of strings with replacements. It is safe for concurrent use
 by multiple goroutines.
 NewReplacer returns a new Replacer from a list of old, new string pairs. Replacements
 are performed in the order they appear in the target string, without overlapping
 matches.
 Replace returns a copy of s with all replacements performed.
 WriteString writes s to w with all replacements performed.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Replacer结构体使用替换字符替换一系列字符串，并行安全，Replacer结构体私有成员r replacer接口
 定义如下：
 // replacer is the interface that a replacement algorithm needs to implement.
 type replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
 }
 1. NewReplacer函数根据一个新旧字符串替换对返回一个新的Replacer结构体，替换按照目标字符串出现顺序
 即新旧新旧新旧...方式进行，且不会重复进行匹配；
 2. Replace方法返回参数字符串s按照Replacer定义的新旧字符串对进行替换后的新的字符串；
 3. WriteString方法把参数字符串s写入参数w，并执行所有替换动作。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
	r.WriteString(os.Stdout, "This is <b>HTML</b>!\n")
}
