/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.Command
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Command(name string, arg ...string) *Cmd
 ------------------------------------------------------------------------------------
 **Description:
 Command returns the Cmd struct to execute the named program with the given arguments.
 It sets only the Path and Args in the returned structure.
 If name contains no path separators, Command uses LookPath to resolve name to a
 complete path if possible. Otherwise it uses name directly as Path.
 The returned Cmd's Args field is constructed from the command name followed by the
 elements of arg, so arg should not include the command name itself. For example,
 Command("echo", "hello"). Args[0] is always name, not the possibly resolved Path.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Command函数用来初始化一个Cmd指针，Path和Args按参数初始化，其他字段执行默认初始化。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	Test2()
}
func Test2() {
	cmd := exec.Command("prog")
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
	)
	if err := cmd.Run(); err != nil {
		println("1", err)
		log.Fatal(err)
	}
}
