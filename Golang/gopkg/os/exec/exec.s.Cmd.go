/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.Cmd
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Cmd struct {
     // Path is the path of the command to run.
     //
     // This is the only field that must be set to a non-zero
     // value. If Path is relative, it is evaluated relative
     // to Dir.
     Path string

     // Args holds command line arguments, including the command as Args[0].
     // If the Args field is empty or nil, Run uses {Path}.
     //
     // In typical use, both Path and Args are set by calling Command.
     Args []string

     // Env specifies the environment of the process.
     // Each entry is of the form "key=value".
     // If Env is nil, the new process uses the current process's
     // environment.
     // If Env contains duplicate environment keys, only the last
     // value in the slice for each duplicate key is used.
     Env []string

     // Dir specifies the working directory of the command.
     // If Dir is the empty string, Run runs the command in the
     // calling process's current directory.
     Dir string

     // Stdin specifies the process's standard input.
     //
     // If Stdin is nil, the process reads from the null device (os.DevNull).
     //
     // If Stdin is an *os.File, the process's standard input is connected
     // directly to that file.
     //
     // Otherwise, during the execution of the command a separate
     // goroutine reads from Stdin and delivers that data to the command
     // over a pipe. In this case, Wait does not complete until the goroutine
     // stops copying, either because it has reached the end of Stdin
     // (EOF or a read error) or because writing to the pipe returned an error.
     Stdin io.Reader

     // Stdout and Stderr specify the process's standard output and error.
     //
     // If either is nil, Run connects the corresponding file descriptor
     // to the null device (os.DevNull).
     //
     // If either is an *os.File, the corresponding output from the process
     // is connected directly to that file.
     //
     // Otherwise, during the execution of the command a separate goroutine
     // reads from the process over a pipe and delivers that data to the
     // corresponding Writer. In this case, Wait does not complete until the
     // goroutine reaches EOF or encounters an error.
     //
     // If Stdout and Stderr are the same writer, and have a type that can
     // be compared with ==, at most one goroutine at a time will call Write.
     Stdout io.Writer
     Stderr io.Writer

     // ExtraFiles specifies additional open files to be inherited by the
     // new process. It does not include standard input, standard output, or
     // standard error. If non-nil, entry i becomes file descriptor 3+i.
     //
     // ExtraFiles is not supported on Windows.
     ExtraFiles []*os.File

     // SysProcAttr holds optional, operating system-specific attributes.
     // Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
     SysProcAttr *syscall.SysProcAttr

     // Process is the underlying process, once started.
     Process *os.Process

     // ProcessState contains information about an exited process,
     // available after a call to Wait or Run.
     ProcessState *os.ProcessState
     // contains filtered or unexported fields
 }
 func Command(name string, arg ...string) *Cmd
 func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
 func (c *Cmd) CombinedOutput() ([]byte, error)
 func (c *Cmd) Output() ([]byte, error)
 func (c *Cmd) Run() error
 func (c *Cmd) Start() error
 func (c *Cmd) StderrPipe() (io.ReadCloser, error)
 func (c *Cmd) StdinPipe() (io.WriteCloser, error)
 func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
 func (c *Cmd) Wait() error
 ------------------------------------------------------------------------------------
 **Description:
 Cmd represents an external command being prepared or run.
 A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.
 Command returns the Cmd struct to execute the named program with the given arguments.
 It sets only the Path and Args in the returned structure.
 If name contains no path separators, Command uses LookPath to resolve name to a
 complete path if possible. Otherwise it uses name directly as Path.
 The returned Cmd's Args field is constructed from the command name followed by the
 elements of arg, so arg should not include the command name itself. For example,
 Command("echo", "hello"). Args[0] is always name, not the possibly resolved Path.
 CommandContext is like Command but includes a context.
 The provided context is used to kill the process (by calling os.Process.Kill) if the
 context becomes done before the command completes on its own.
 CombinedOutput runs the command and returns its combined standard output and standard
 error.
 Output runs the command and returns its standard output. Any returned error will
 usually be of type *ExitError. If c.Stderr was nil, Output populates ExitError.Stderr.
 Run starts the specified command and waits for it to complete.
 The returned error is nil if the command runs, has no problems copying stdin, stdout,
 and stderr, and exits with a zero exit status.
 If the command starts but does not complete successfully, the error is of type
 *ExitError. Other error types may be returned for other situations.
 If the calling goroutine has locked the operating system thread with
 runtime.LockOSThread and modified any inheritable OS-level thread state
 (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's
 thread state.
 Start starts the specified command but does not wait for it to complete.
 The Wait method will return the exit code and release associated resources once the
 command exits.
 StderrPipe returns a pipe that will be connected to the command's standard error when
 the command starts.
 Wait will close the pipe after seeing the command exit, so most callers need not
 close the pipe themselves; however, an implication is that it is incorrect to call
 Wait before all reads from the pipe have completed. For the same reason, it is
 incorrect to use Run when using StderrPipe. See the StdoutPipe example for idiomatic
 usage.
 StdinPipe returns a pipe that will be connected to the command's standard input when
 the command starts. The pipe will be closed automatically after Wait sees the command
 exit. A caller need only call Close to force the pipe to close sooner. For example,
 if the command being run will not exit until standard input is closed, the caller
 must close the pipe.
 StdoutPipe returns a pipe that will be connected to the command's standard output
 when the command starts.
 Wait will close the pipe after seeing the command exit, so most callers need not
 close the pipe themselves; however, an implication is that it is incorrect to call
 Wait before all reads from the pipe have completed. For the same reason, it is
 incorrect to call Run when using StdoutPipe. See the example for idiomatic usage.
 Wait waits for the command to exit and waits for any copying to stdin or copying
 from stdout or stderr to complete.
 The command must have been started by Start.
 The returned error is nil if the command runs, has no problems copying stdin, stdout,
 and stderr, and exits with a zero exit status.
 If the command fails to run or doesn't complete successfully, the error is of type
 *ExitError. Other error types may be returned for I/O problems.
 If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits for
 the respective I/O loop copying to or from the process to complete.
 Wait releases any resources associated with the Cmd.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. CombinedOutput方法执行*Cmd中的命令，把执行结果和错误合并到byte切片中返回；
 2. Output方法执行Cmd中的命令，并返回执行结果；
 3. Run执行*Cmd中的命令，并等待命令执行完成。若执行不成功，返回error；
 4. Start方法执行*Cmd中的命令，它只是让命令开始执行，并不会等待命令执行完；
 5. StderrPipe返回一个io.ReadCloser,用于连接命令启动时错误标准输出的管道，命令结束时，
 管道会自动关闭；
 6. StdinPipe方法用于连接到命令启动时标准输入的管道
 7. StdoutPipe方法用于连接到命令启动时标准输出的管道，命令结束时，管道会自动关闭
 8. Wait方法等待*Cmd中的已开始执行的命令执行完成。
 9. Cmd结构体的各个方法使用还没有搞得非常清晰，需要后续实践后更新理解。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func TestCombinedOutput() {
	arg := []string{"./_iofiles/Hello", "./_iofiles/World"}
	cmd := exec.Command("mv", arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err) //Error: exit status 1
	}
	//test in ArchLinux
	//当前目录中必须没有名为Hello的文件或文件夹
	fmt.Printf("The output is: %s\n", output) //The output is: mv: 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
}

func TestOutput() {
	arg := []string{"Hello", "World!"}
	cmd := exec.Command("echo", arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output) //The output is: Hello World!
}
func TestRun() {
	arg := []string{"Hello", "World!"}
	cmd := exec.Command("echo", arg...)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: Hello World!
}
func TestStart() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Waiting for command to finish...\n")
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)
}
func TestStderrPipe() {
	arg := []string{"Hello", "World!"}
	cmd := exec.Command("mv", arg...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	output, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err = cmd.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err) //Error: exit status 1
	}
	//test in ArchLinux
	//当前目录中必须没有名为Hello的文件或文件夹
	fmt.Printf("The mv command error is: %s\n", output) //The mv command error is: mv: 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
}
func TestStdinPipe() {
	var output bytes.Buffer
	cmd := exec.Command("cat")
	cmd.Stdout = &output
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	stdin.Write([]byte("Hello World!"))
	stdin.Close()
	if err = cmd.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: Hello World!
}
func TestStdoutPipe() {
	cmd := exec.Command("echo", "Hello World!")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	output, _ := ioutil.ReadAll(stdout)
	if err = cmd.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output) //The output is: Hello World!
}

func main() {
	fmt.Println("------CombinedOutput----------")
	TestCombinedOutput()
	TestCombinedOutput()
	fmt.Println("------Output----------")
	TestOutput()
	fmt.Println("------Run----------")
	TestRun()
	fmt.Println("------Start----------")
	TestStart()
	fmt.Println("------StderrPipe----------")
	TestStderrPipe()

	fmt.Println("------StdinPipe----------")
	TestStdinPipe()
	fmt.Println("------StdoutPipe----------")
	TestStdoutPipe()
	fmt.Println("------Wait----------")
	TestWait()
}
func TestWait() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Waiting for command to finish...\n")
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)
}
