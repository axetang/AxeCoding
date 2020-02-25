/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.ProcessState
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ProcessState struct {
     // contains filtered or unexported fields
 }
 func (p *ProcessState) Exited() bool
 func (p *ProcessState) Pid() int
 func (p *ProcessState) String() string
 func (p *ProcessState) Success() bool
 func (p *ProcessState) Sys() interface{}
 func (p *ProcessState) SysUsage() interface{}
 func (p *ProcessState) SystemTime() time.Duration
 func (p *ProcessState) UserTime() time.Duration
 ------------------------------------------------------------------------------------
 **Description:
 ProcessState stores information about a process, as reported by Wait.
 Exited reports whether the program has exited.
 Pid returns the process id of the exited process.
 Success reports whether the program exited successfully, such as with exit status 0
 on Unix.
 Sys returns system-dependent exit information about the process. Convert it to the
 appropriate underlying type, such as syscall.WaitStatus on Unix, to access its
 contents.
 SysUsage returns system-dependent resource usage information about the exited
 process. Convert it to the appropriate underlying type, such as *syscall.Rusage on
 Unix, to access its contents. (On Unix, *syscall.Rusage matches struct rusage as
 defined in the getrusage(2) manual page.)
 SystemTime returns the system CPU time of the exited process and its children.
 UserTime returns the user CPU time of the exited process and its children.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ProcessState结构体存储一个进程的所有状态信息；
 2. Exited方法判断进程状态是不是已经退出；
 3. Pid方法返回已经退出的进程的pid；
 4. Success方法判断进程是否成功退出，即exit status 0；
 5. Sys方法返回当前进程的退出信息；
 6. SysUsage方法返回已经退出进程的系统资源使用信息；
 7. SystemTime返回已经退出的进程及其子进程占用的系统CPU时间；
 8. UserTime返回已经退出进程及其子进程的用户CPU时间。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func TestExited() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process has been exited?: %t\n", ps.Exited())
}

func TestPid() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process pid is: %d\n", ps.Pid())
}

func TestString() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process stat: %s\n", ps.String())
}
func TestSuccess() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process has been exited successfully?: %t\n", ps.Success())
}
func TestSys() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process exit status: %+v\n", ps.Sys())
}

func TestSystemTime() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process system cpu time: %+v\n", ps.SystemTime())
}
func TestSysUsage() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process resource usage: %+v\n", ps.SysUsage())
}
func TestUserTime() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}
	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ps, _ := p.Wait()
	fmt.Printf("Process user cpu time: %+v\n", ps.UserTime())
}

func main() {
	fmt.Println("-------Exited--------")
	TestExited()
	fmt.Println("-------Pid--------")
	TestPid()
	fmt.Println("-------String--------")
	TestString()
	fmt.Println("-------Success--------")
	TestSuccess()
	fmt.Println("-------Sys--------")
	TestSys()
	fmt.Println("-------SystemTime--------")
	TestSystemTime()
	fmt.Println("-------SysUsage--------")
	TestSysUsage()
	fmt.Println("-------UserTime--------")
	TestUserTime()
}
