/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Process
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Process struct {
     Pid int
     // contains filtered or unexported fields
 }
 func (p *Process) Kill() error
 func (p *Process) Release() error
 func (p *Process) Signal(sig Signal) error
 func (p *Process) Wait() (*ProcessState, error)
 ------------------------------------------------------------------------------------
 **Description:
 Process stores the information about a process created by StartProcess.
 Kill causes the Process to exit immediately. Kill does not wait until the Process has
 actually exited. This only kills the Process itself, not any other processes it may
 have started.
 Release releases any resources associated with the Process p, rendering it unusable
 in the future. Release only needs to be called if Wait is not.
 Signal sends a signal to the Process. Sending Interrupt on Windows is not implemented.
 Wait waits for the Process to exit, and then returns a ProcessState describing its
 status and an error, if any. Wait releases any resources associated with the Process.
 On most operating systems, the Process must be a child of the current process or an
 error will be returned.
------------------------------------------------------------------------------------
 **要点总结:
 1. Process结构体存储了一个使用StartProcess创建的进程的所有信息，其中只有成员Pid是导出的公开成员；
 2. Kill方法立即杀死进程p并退出；
 3. Release方法释放与进程p有关的的所有资源；
 4. Signal方法向进程p发送信号sig Signal；
 5. Wait方法等待进程完成并退出，并返回一个*ProcessState来描述该进程的状态，Wait方法会释放与该进程
 有关的所有资源。
 6. Wait是一个非常重要的方法，要获取一个进程p的状态ProcessStat，就需要调用这个wait方法，获得返回值
 *ProcessState，然后调用ProcessState的方法，如Exited、String、Pid等等来获取进程具体状态信息。
 *************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func TestKill() {
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
	if err := p.Kill(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func TestRelease() {
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
	if err := p.Release(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func TestWait() {
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
	fmt.Printf("Process stat: %+v\n", ps)
}
func TestSignal() {
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
	if err := p.Signal(os.Kill); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	fmt.Println("----Kill-----")
	TestKill()
	fmt.Println("----Release-----")
	TestRelease()
	fmt.Println("----Wait-----")
	TestWait()
	fmt.Println("----Signal-----")
	TestSignal()
}
