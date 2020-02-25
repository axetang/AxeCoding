/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.ProcAttr
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ProcAttr struct {
     // If Dir is non-empty, the child changes into the directory before
     // creating the process.
     Dir string
     // If Env is non-nil, it gives the environment variables for the
     // new process in the form returned by Environ.
     // If it is nil, the result of Environ will be used.
     Env []string
     // Files specifies the open files inherited by the new process. The
     // first three entries correspond to standard input, standard output, and
     // standard error. An implementation may support additional entries,
     // depending on the underlying operating system. A nil entry corresponds
     // to that file being closed when the process starts.
     Files []*File

     // Operating system-specific process creation attributes.
     // Note that setting this field means that your program
     // may not execute properly or even compile on some
     // operating systems.
     Sys *syscall.SysProcAttr
 }
------------------------------------------------------------------------------------
 **Description:
 ProcAttr holds the attributes that will be applied to a new process started by
 StartProcess.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ProcAttr结构体包含四个成员:Dir,Env,Files,Sys;
 2. Sys成员不用设置，是操作系统指定的进程创建属性；
 3. ProcAttr结构体包含了使用StartProcss函数创建进程所需的所有属性。
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
	p, err := os.StartProcess("/bin/ls", []string{"/bin/ls", "-l"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process info: %+v\n", p)
}
