/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.StartProcess
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
 ------------------------------------------------------------------------------------
 **Description:
 StartProcess starts a new process with the program, arguments and attributes
 specified by name, argv and attr. The argv slice will become os.Args in the new
 process, so it normally starts with the program name.
 If the calling goroutine has locked the operating system thread with
 runtime.LockOSThread and modified any inheritable OS-level thread state (for example,
 Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.
 StartProcess is a low-level interface. The os/exec package provides higher-level
 interfaces.
 If there is an error, it will be of type *PathError.
------------------------------------------------------------------------------------
 **要点总结:
 1. StartProcess方法按照指定参数启动一个进程；
 2. name代表要创建进程的带路径的完整名称；
 3. argv []string这个slice成为启动的这个进程的os.Args，当然是包含name的，name及时argv[0]；
 4. attr代表指定的进程属性*ProcAttr；
 5. 返回创建进程的指针*Process，如果有错误发生，返回PathError错误。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process info: %+v\n", p)
}
