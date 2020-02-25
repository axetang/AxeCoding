/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.ExitError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ExitError struct {
     *os.ProcessState

     // Stderr holds a subset of the standard error output from the
     // Cmd.Output method if standard error was not otherwise being
     // collected.
     //
     // If the error output is long, Stderr may contain only a prefix
     // and suffix of the output, with the middle replaced with
     // text about the number of omitted bytes.
     //
     // Stderr is provided for debugging, for inclusion in error messages.
     // Users with other needs should redirect Cmd.Stderr as needed.
     Stderr []byte
 }
 func (e *ExitError) Error() string
------------------------------------------------------------------------------------
 **Description:
 An ExitError reports an unsuccessful exit by a command.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ExitError存储一个命令不成功的退出的错误信息；
 2. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("mv", "Hello World!")
	cmd.Run()
	exitError := exec.ExitError{cmd.ProcessState, []byte("")}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", exitError.Error()) //The output is: exit status 1
}
